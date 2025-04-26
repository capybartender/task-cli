package main

import (
	"fmt"
	"os"

	"github.com/capybartender/task-cli/internal/validators"
)

// requirements https://roadmap.sh/projects/task-tracker

func hello() string {
	return "Hello there!"
}



func main() {
	args := os.Args[1:]

	isValid, err := validators.ValidateArgs(args)

	if !isValid {
		fmt.Println(err.Error())
		return
	}

	command := args[0]

	fmt.Println(command)
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
