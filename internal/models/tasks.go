package models

import (
	"fmt"
	"slices"
	"time"
)

type Task struct {
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskList []Task

func (t *TaskList) MarkDone(id int) error {
	return changeStatus(t, id, Done)
}

func validateIndex(t *TaskList, index int) error {
	if len(*t)-1 > index {
		return nil
	}
	return fmt.Errorf("task with id %d does not exist", index)
}

func changeStatus(t *TaskList, index int, newStatus TaskStatus) error {
	if err := validateIndex(t, index); err != nil {
		return err
	}

	(*t)[index].Status = newStatus
	(*t)[index].UpdatedAt = time.Now()
	return nil
}

func (t *TaskList) MarkInProgress(id int) error {
	return changeStatus(t, id-1, InProgress)
}

func (t *TaskList) Delete(id int) error {
	index := id - 1
	if err := validateIndex(t, index); err != nil {
		return err
	}
	if index == 0 {
		*t = (*t)[1:]
		return nil
	}
	if index == len(*t)-1 {
		*t = (*t)[:index]
		return nil
	}
	(*t) = slices.Delete((*t), index-1, index+1)
	return nil
}

func (t *TaskList) Update(id int, newDescription string) error {
	index := id - 1
	if err := validateIndex(t, index); err != nil {
		return err
	}

	(*t)[index].Description = newDescription
	(*t)[index].UpdatedAt = time.Now()
	return nil
}

func (t *TaskList) Add(newTask string) error {
	*t = append(*t, Task{
		Description: newTask,
		Status:      ToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	return nil
}

func (t *TaskList) ListByStatus(status TaskStatus) []string {
	var out []string
	for index, task := range *t {
		if task.Status == status {
			out = append(out, toString(index+1, &task))
		}
	}
	return out
}

func (t *TaskList) ListAll() []string {
	out := []string{}
	for index, task := range *t {
		out = append(out, toString(index+1, &task))
	}
	return out
}

func toString(id int, task *Task) string {
	return fmt.Sprintf("%02d: %s", id, *&task.Description)
}
