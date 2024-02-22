// task.go

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Task represents a unit of work to be executed by the task scheduler.
type Task struct {
	ID       int           // Unique identifier for the task
	Name     string        // Name or description of the task
	Function func() error  // Function to execute as part of the task
	Logger   logrus.FieldLogger // Logger interface
}



// Execute executes the task function.
func (t *Task) Execute() {
    fmt.Printf("Executing task: %s\n", t.Name)
    t.Function() // Execute the task function
}