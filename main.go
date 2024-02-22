// main.go

package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)



func init() {
	// Initialize a new logger instance
	logger = logrus.New()

	// Set logger output to stdout
	logger.SetOutput(os.Stdout)

	// Set logger level to debug
	logger.SetLevel(logrus.DebugLevel)
}

// SampleTaskFunction is a sample function to be used as a task function.
func SampleTaskFunction() error {
	// Log task execution start
	logger.Debugf("Executing task function")

	// Simulate some task execution
	fmt.Println("Executing sample task function")

	// Log task execution completion
	logger.Debugf("Task function execution completed")

	return nil
}

// Main function
func main() {
	// Specify the maximum number of concurrent tasks
	concurrentLimit := 2

	// Instantiate a task scheduler with concurrent limit
	scheduler := TaskScheduler{
		ConcurrentLimit: concurrentLimit,
	}

	// Create a new task
	task1 := Task{
		ID:       1,
		Name:     "Sample Task 1",
		Function: SampleTaskFunction,
	}

	// Add the task to the scheduler
	scheduler.AddTask(&task1)

	// Create another task
	task2 := Task{
		ID:       2,
		Name:     "Sample Task 2",
		Function: SampleTaskFunction,
	}

	// Add the second task to the scheduler
	scheduler.AddTask(&task2)

	// Execute the tasks concurrently
	scheduler.executeTasks(scheduler.tasks)

}
