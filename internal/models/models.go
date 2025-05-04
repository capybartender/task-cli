package models

import "fmt"

type Command struct {
	Command string
	Args    []string
}

type Output struct {
	Value []string
}

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

func ToTaskStatus(s string) (TaskStatus, error) {
	if status, ok := statusValue[s]; ok {
		return status, nil
	}
	return -1, fmt.Errorf("invalid task status: %s", s)
}
