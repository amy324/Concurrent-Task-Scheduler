// main.go

package main

import "fmt"

func main() {
    // Create a new task
    task := Task{
        ID:       1,
        Name:     "Sample Task",
        Function: SampleTaskFunction,
    }
    
    // Print task details
    fmt.Println("Task ID:", task.ID)
    fmt.Println("Task Name:", task.Name)
    
    // Execute the task function
    err := task.Function()
    if err != nil {
        fmt.Println("Error executing task:", err)
    }
}
