package taskprocessor

import (
	"fmt"
	"strings"

	"github.com/capybartender/task-cli/internal/models"
)

type TaskProcessor struct {
	file string
}

const defaultFileName = "tasks.json"

func (t *TaskProcessor) Execute(c *models.Command) (*models.Output, error) {
	if c == nil {
		return nil, fmt.Errorf("Command is nil")		
	}

	if strings.TrimSpace(t.file) == "" {
		t.file = defaultFileName		
	}

	tasks, err := load(t.file)	
	if err != nil {
		return nil, err
	}

	var output *models.Output = nil

	switch c.Command {
	case "list":
		err := tasks.List(c.Args)
		if err == nil {
			output = &models.Output{ Value: tasks.GetTaskStrings()}			
		}
		return output, err
	case "add":
		err = tasks.Add(c.Args)
		break
	case "update":
		err = tasks.Update(c.Args)
		break
	case "delete":
		err = tasks.Delete(c.Args)
		break
	case "mark-in-progress":
		err = tasks.MarkInProgress(c.Args)
		break
	case "mark-done":
		err = tasks.MarkDone(c.Args)
		break
	default:
		return nil, fmt.Errorf("Unknown command: %s", c.Command)
	}
	
	if err != nil {
		return nil, err
	}

	err = save(t.file, tasks) // Save the tasks after executing the command
	
	if err != nil {
		return nil, err
	}
}

func load(file string) (*models.TaskList, error) {
	// Implement the logic to load tasks from the file
	return nil, nil
}

func save(file string, tasks *models.TaskList) error {
	// Implement the logic to save tasks to the file
	return nil
}