package models

import (
	"time"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskList []Task

func (t *TaskList) MarkDone(args []string) error {
	panic("unimplemented")
}

func (t *TaskList) MarkInProgress(args []string) error {
	panic("unimplemented")
}

func (t *TaskList) Delete(args []string) error {
	panic("unimplemented")
}

func (t *TaskList) Update(args []string) error {
	panic("unimplemented")
}

func (t *TaskList) Add(args []string) error {
	panic("unimplemented")
}

func (t *TaskList) List(args []string) error {
	panic("unimplemented")
}

func (t *TaskList) GetTaskStrings() []string {
	taskStrings := []string{}
	for _, task := range *t {
		taskStrings = append(taskStrings, task.Description)
	}
	return taskStrings
}
