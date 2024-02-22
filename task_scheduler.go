// task_scheduler.go

package main

import (
    "sync"
)

// TaskScheduler represents a concurrent task scheduler.
type TaskScheduler struct {
    tasks []*Task        // List of tasks to be executed
    mutex sync.Mutex     // Mutex for synchronization
}

// AddTask adds a new task to the task scheduler.
func (ts *TaskScheduler) AddTask(task *Task) {
    // Lock the mutex to ensure exclusive access to the task list
    ts.mutex.Lock()
    defer ts.mutex.Unlock()

    // Add the task to the task list
    ts.tasks = append(ts.tasks, task)
}