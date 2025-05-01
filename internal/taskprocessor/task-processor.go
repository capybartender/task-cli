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
			var status models.TaskStatus
			status, err = models.ToTaskStatus(c.Args[0])
			if err != nil {
				return nil, err
			}
			taskList = tasks.ListByStatus(status)
		}
		output = &models.Output{Value: taskList}
		return output, nil
	case "add":
		var result string
		result, err = tasks.Add(c.Args[0])
		if err != nil {
			return nil, err
		}
		output = &models.Output{Value: []string{result}}
	case "update":
		// var id int
		// if id, err = strconv.Atoi(c.Args[0]); err == nil {
		// 	err = tasks.Update(id, c.Args[1])
		// }

		err = parseArgAsIdAndExecuteFunction(c.Args[0], partialUpdate(tasks, c.Args[1]))
	case "delete":
		// var id int
		// if id, err = strconv.Atoi(c.Args[0]); err == nil {
		// 	err = tasks.Delete(id)
		// }
		err = parseArgAsIdAndExecuteFunction(c.Args[0], tasks.Delete)
	case "mark-in-progress":
		// var id int
		// if id, err = strconv.Atoi(c.Args[0]); err == nil {
		// 	err = tasks.MarkInProgress(id)
		// }
		err = parseArgAsIdAndExecuteFunction(c.Args[0], tasks.MarkInProgress)
	case "mark-done":
		// var id int
		// if id, err = strconv.Atoi(c.Args[0]); err == nil {
		// 	err = tasks.MarkDone(id)
		// }

		err = parseArgAsIdAndExecuteFunction(c.Args[0], tasks.MarkDone)
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

	return output, nil
}

func partialUpdate(t *models.TaskList, text string) func(int) error {
	return func(id int) error {
		return t.Update(id, text)
	}
}

func parseArgAsIdAndExecuteFunction(idString string, f func(int) error) error {
	if id, err := strconv.Atoi(idString); err == nil {
		return f(id)
	}
	return nil
}
