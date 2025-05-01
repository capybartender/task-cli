package taskprocessor

import (
	"fmt"
	"strconv"

	"github.com/capybartender/task-cli/internal/models"
	"github.com/capybartender/task-cli/internal/storage"
)

type TaskProcessor struct{}

var taskStorage = storage.Init("")

func (t *TaskProcessor) Execute(c *models.Command) (*models.Output, error) {
	if c == nil {
		return nil, fmt.Errorf("command is nil")
	}

	tasks, err := taskStorage.Load()
	if err != nil {
		return nil, err
	}

	var output *models.Output = nil

	switch c.Command {
	case "list":
		var taskList []string
		if len(c.Args) == 0 {
			taskList = tasks.ListAll()
		} else {
			status, err := models.ToTaskStatus(c.Args[0])
			if err != nil {
				return nil, err
			}
			taskList = tasks.ListByStatus(status)
		}
		output = &models.Output{Value: taskList}
		return output, nil
	case "add":
		err = tasks.Add(c.Args[0])
	case "update":
		if id, err := strconv.Atoi(c.Args[0]); err == nil {
			err = tasks.Update(id, c.Args[1])
		}
	case "delete":
		if id, err := strconv.Atoi(c.Args[0]); err == nil {
			err = tasks.Delete(id)
		}
	case "mark-in-progress":
		if id, err := strconv.Atoi(c.Args[0]); err == nil {
			err = tasks.MarkInProgress(id)
		}
	case "mark-done":
		if id, err := strconv.Atoi(c.Args[0]); err == nil {
			err = tasks.MarkDone(id)
		}
	default:
		return nil, fmt.Errorf("unknown command: %s", c.Command)
	}

	if err != nil {
		return nil, err
	}

	err = taskStorage.Save(tasks) // Save the tasks after executing the command

	if err != nil {
		return nil, err
	}

	return &models.Output{}, nil
}
