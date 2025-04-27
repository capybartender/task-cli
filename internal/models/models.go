package models

import (
	"time"
)

type TaskStatus int

const (
	ToDo TaskStatus = iota
	InProgress
	Done
)

var statusName = map[TaskStatus]string{
	ToDo:       "todo",
	InProgress: "in-progress",
	Done:       "done",
}

var statusValue = map[string]TaskStatus{
	"todo":        ToDo,
	"in-progress": InProgress,
	"done":        Done,
}

func (ts TaskStatus) String() string {
	return statusName[ts]
}

func IsValidState(value string) bool {
	if _, ok := statusValue[value]; ok {
		return true
	}
	return false
}

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type Output struct {
	Value []string `json:"value"`
}
