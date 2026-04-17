// task.go
package main

import "time"

type Task struct {
	Id 			int
	Description string
	Status 		string
	CreatedAt  	time.Time
	UpdatedAt	time.Time
}

const (
	StatusTodo			= "todo"
	StatusInProgress 	= "in-progress"
	StatusDone			= "done"
)