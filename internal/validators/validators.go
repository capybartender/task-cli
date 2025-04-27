package validators

import (
	"errors"
	"slices"
	"strconv"

	"github.com/capybartender/task-cli/internal/models"
)

func ValidateArgs(args []string) (*models.Command, error) {
	argsLength := len(args)

	if argsLength == 0 {
		return nil, errors.New("This command requires at least one argument.")
	}

	command := args[0]

	oneArgsCommands := []string{"list"}
	twoArgsCommands := []string{"add", "delete", "mark-in-progress", "mark-done", "list"}
	threeArgsCommands := []string{"update"}

	if !slices.Contains(oneArgsCommands, command) && !slices.Contains(twoArgsCommands, command) && !slices.Contains(threeArgsCommands, command) {
		return nil, errors.New("Unknown action.")
	}

	switch argsLength {
	case 0:
		return nil, errors.New("This command requires at least one argument.")
	case 1:
		if !slices.Contains(oneArgsCommands, command) {
			return nil, errors.New("Arguments missing.")
		}
		break
	case 2:
		if !slices.Contains(twoArgsCommands, command) {
			return nil, errors.New("Arguments missing.")
		} else {
			if command != "add" {
				if command == "list" {
					if !isTaskState(args[1]) {
						return nil, errors.New("Second argument is not a valid task state")
					}
				} else if !isNumber(args[1]) {
					return nil, errors.New("Second argument is not an ID")
				}
			}
		}
		break
	case 3:
		if !slices.Contains(threeArgsCommands, command) {
			return nil, errors.New("Arguments missing.")
		} else {
			if !isNumber(args[1]) {
				return nil, errors.New("Second argument is not an ID")
			}
		}
	default:
		return nil, errors.New("Wrong number of arguments")
	}

	return &models.Command{Command: command, Args: args[1:]}, nil
}

func isNumber(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}
	return false
}

func isTaskState(value string) bool {
	return models.IsValidState(value)
}
