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
	ConcurrentLimit int            // Maximum number of concurrent tasks allowed
	Logger          logrus.FieldLogger // Logger interface
}



// Define a mutex for synchronization
var mutex sync.Mutex
// AddTask adds a new task to the task scheduler.
func (ts *TaskScheduler) AddTask(task *Task) {
	// Lock the mutex to ensure exclusive access to the task list
	mutex.Lock()
	defer mutex.Unlock()

	// Set the logger for the task
	task.Logger = ts.Logger

	// Add the task to the task list
	ts.tasks = append(ts.tasks, task)
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
			logger.Debugf("Executing task: %s", t.Name)
			t.Execute() // Simply call t.Execute() without attempting to use its result
		}(task)
	}
	wg.Wait()
}
