// task_scheduler.go

package main

import (
	
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	// Initialize a new logger instance
	logger = logrus.New()

	// Set logger output to stdout
	logger.SetOutput(os.Stdout)

	// Set logger level to debug
	logger.SetLevel(logrus.DebugLevel)
}

// TaskScheduler represents a concurrent task scheduler.
type TaskScheduler struct {
	tasks           []*Task
	ConcurrentLimit int // Maximum number of concurrent tasks allowed
}

// Define a mutex for synchronization
var mutex sync.Mutex

// AddTask adds a new task to the task scheduler.
func (ts *TaskScheduler) AddTask(task *Task) {
	// Lock the mutex to ensure exclusive access to the task list
	mutex.Lock()
	defer mutex.Unlock()

	// Add the task to the task list
	ts.tasks = append(ts.tasks, task)
}

// ExecuteTasks executes tasks concurrently with a limit on concurrent execution.
func (ts *TaskScheduler) ExecuteTasks() {
	// Use a buffered channel to enforce the concurrent limit
	sem := make(chan struct{}, ts.ConcurrentLimit)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Iterate over each task and execute it
	for _, task := range ts.tasks {
		// Increment the WaitGroup counter
		wg.Add(1)

		// Acquire a slot from the semaphore
		sem <- struct{}{}

		// Execute the task in a separate goroutine
		go func(t *Task) {
			defer func() {
				// Release the slot back to the semaphore
				<-sem
				// Decrement the WaitGroup counter when the goroutine completes
				wg.Done()
			}()

			// Log task execution start
			logger.Debugf("Executing task: %s", t.Name)

			// Execute the task function
			if err := t.Function(); err != nil {
				// Log error if task execution fails
				logger.Errorf("Error executing task %s: %v", t.Name, err)
			}

			// Log task execution completion
			logger.Debugf("Task execution completed: %s", t.Name)
		}(task)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Run starts executing tasks from the task scheduler.
func (ts *TaskScheduler) Run() {
	for {
		// Lock the mutex to ensure exclusive access to the task list
		mutex.Lock()

		// Check if there are any tasks in the scheduler
		if len(ts.tasks) == 0 {
			// If no tasks are available, release the mutex and wait for new tasks
			mutex.Unlock()
			time.Sleep(1 * time.Second) // Adjust the sleep duration as needed
			continue
		}

		// Check if the number of concurrent tasks has reached the limit
		if len(ts.tasks) > ts.ConcurrentLimit {
			// If the number of tasks exceeds the limit, execute only up to the limit
			ts.executeTasks(ts.tasks[:ts.ConcurrentLimit])
			ts.tasks = ts.tasks[ts.ConcurrentLimit:]
		} else {
			// If the number of tasks is within the limit, execute all tasks
			ts.executeTasks(ts.tasks)
			ts.tasks = nil
		}

		// Release the mutex after processing tasks
		mutex.Unlock()
	}
}

// executeTasks executes the given list of tasks concurrently.
func (ts *TaskScheduler) executeTasks(tasks []*Task) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t *Task) {
			defer wg.Done()
			t.Execute() // Execute the task
		}(task)
	}
	wg.Wait()
}
