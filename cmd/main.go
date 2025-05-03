package main

import (
	"fmt"
	"os"

	"github.com/capybartender/task-cli/internal/models"
	"github.com/capybartender/task-cli/internal/taskprocessor"
	"github.com/capybartender/task-cli/internal/validators"
)

// requirements https://roadmap.sh/projects/task-tracker

func hello() string {
	return "Hello there!"
}

func main() {
	args := os.Args[1:]

	command, err := validators.ValidateArgs(args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	taskprocessor := taskprocessor.TaskProcessor{}

	var output *models.Output
	output, err = taskprocessor.Execute(command)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if output != nil && output.Value != nil && len(output.Value) > 0 {
		fmt.Println("Tasks:")
		for _, task := range output.Value {
			fmt.Println(task)
		}
	}
}

/*
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
*/
