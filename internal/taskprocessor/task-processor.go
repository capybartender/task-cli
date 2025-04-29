package taskprocessor

import (
	"fmt"

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
		err := tasks.List(c.Args)
		if err == nil {
			output = &models.Output{Value: tasks.GetTaskStrings()}
		}
		return output, err
	case "add":
		err = tasks.Add(c.Args)
	case "update":
		err = tasks.Update(c.Args)
	case "delete":
		err = tasks.Delete(c.Args)
	case "mark-in-progress":
		err = tasks.MarkInProgress(c.Args)
	case "mark-done":
		err = tasks.MarkDone(c.Args)
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
