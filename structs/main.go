package main

import "fmt"

/*
PATTERN: DATA + BEHAVIOR = DOMAIN MODEL

A struct groups related data.
Methods define what that data can do.
*/

// STRUCT: Represent a task as one meaningful type.
type Task struct {
	Title string
	Done  bool
}

// VALUE RECEIVER: Read the task without modifying it.
func (t Task) Status() string {
	if t.Done {
		return "Completed"
	}

	return "Pending"
}

// POINTER RECEIVER: Modify the original task's state.
func (t *Task) Complete() {
	t.Done = true
}

func main() {
	// Create a task with its initial state.
	task := Task{
		Title: "Learn Go patterns",
		Done:  false,
	}

	// Read the current state.
	fmt.Println(task.Title, ":", task.Status())

	// Update the original task using a pointer receiver.
	task.Complete()

	// The task now has a different state.
	fmt.Println(task.Title, ":", task.Status())
}
