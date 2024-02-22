

# Concurrent Task Scheduler

This is a concurrent task scheduler written in Go. It allows you to execute multiple tasks concurrently while limiting the maximum number of tasks running simultaneously.

## Overview

The program consists of several files:

- `main.go`: Contains the main entry point of the program and demonstrates how to use the task scheduler.
- `task.go`: Defines the `Task` struct, representing a unit of work to be executed.
- `task_scheduler.go`: Defines the `TaskScheduler` struct, which manages the execution of tasks concurrently.
- `task_test.go`: Contains unit tests for the task-related functionality.

## Features

- **Concurrent Execution:** Execute multiple tasks concurrently, maximizing resource utilization.
- **Concurrent Limit:** Limit the maximum number of tasks running simultaneously to prevent resource exhaustion.
- **Logging:** Utilize the Logrus package for structured logging, enabling detailed logging of task execution.


## Methods Used

### Goroutines

Goroutines are lightweight threads managed by the Go runtime, allowing concurrent execution of tasks. They are used in the `executeTasks` method of the `TaskScheduler` struct to execute tasks concurrently.

Example usage:

```go
// executeTasks executes the given list of tasks concurrently.
func (ts *TaskScheduler) executeTasks(tasks []*Task) {
    var wg sync.WaitGroup
    for _, task := range tasks {
        wg.Add(1)
        go func(t *Task) {
            defer wg.Done()
            logger.Debugf("Executing task: %s", t.Name)
            t.Execute()
        }(task)
    }
    wg.Wait()
}
```

### Mutexes

Mutexes are synchronization primitives used to protect shared resources from concurrent access. They are used in the `AddTask` method of the `TaskScheduler` struct to ensure exclusive access to the task list when adding tasks.

Example usage:

```go
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
```

### Pointers

Pointers are used for efficient memory management and passing objects by reference. They are used to pass task objects to methods and functions efficiently without unnecessary copying of data.

Example usage:

```go
// Create a new task
task1 := Task{
    ID:       1,
    Name:     "Sample Task 1",
    Function: SampleTaskFunction,
}

// Add the task to the scheduler
scheduler.AddTask(&task1)
```

### Logrus

Logrus is a structured logger for Go, used for logging within the program. It provides rich logging capabilities, including support for different log levels and log formatting.

Example usage:

```go
import (
    "os"
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
```

## Usage

To use the concurrent task scheduler, follow these steps:

1. Import the package into your Go program.
2. Create tasks using the `Task` struct, specifying the task function and any other relevant details.
3. Add tasks to the scheduler using the `AddTask` method.
4. Run the task scheduler using the `Run` method to start executing tasks concurrently.

Example:

```bash
$ go run .
time="2024-02-22T17:50:19Z" level=debug msg="Executing task: Sample Task 2"
Executing task: Sample Task 2
time="2024-02-22T17:50:19Z" level=debug msg="Executing task: Sample Task 1"
Executing task: Sample Task 1
time="2024-02-22T17:50:19Z" level=debug msg="Executing task function"
Executing sample task function
time="2024-02-22T17:50:19Z" level=debug msg="Executing task function"
Executing sample task function
time="2024-02-22T17:50:19Z" level=debug msg="Task function execution completed"
time="2024-02-22T17:50:19Z" level=debug msg="Task function execution completed"
```



## Testing

The program includes unit tests to verify the correctness of the task-related functionality. These tests ensure that tasks execute as expected and that the task scheduler behaves correctly under various conditions.

To run the tests, execute the following command:

```bash
go test
```
Example:

```bash
$ go test
Executing task: Test Task
time="2024-02-22T17:51:03Z" level=debug msg="Executing task function"
Executing sample task function
time="2024-02-22T17:51:03Z" level=debug msg="Task function execution completed"
PASS
ok      task-scheduler  0.157s
```



## Dependencies

- [Logrus](https://github.com/sirupsen/logrus): A structured logger for Go.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

