// main.go

package main

import (
    "fmt"
)

func main() {
    // Instantiate a task scheduler
    scheduler := TaskScheduler{}

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

    // Print the tasks in the scheduler
    fmt.Println("Tasks in the scheduler:")
    for _, task := range scheduler.tasks {
        fmt.Println("Task ID:", task.ID)
        fmt.Println("Task Name:", task.Name)
    }
}
