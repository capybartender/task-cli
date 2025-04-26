package validators

import (
	"errors"
	"slices"
	"strconv"

	"github.com/capybartender/task-cli/internal/models"
)

func ValidateArgs(args []string) (bool, error) {
	argsLength := len(args)

	if argsLength == 0 {
		return false, errors.New("This command requires at least one argument.")
	}

	command := args[0]

	oneArgsCommands := []string{"list"}
	twoArgsCommands := []string{"add", "delete", "mark-in-progress", "mark-done", "list"}
	threeArgsCommands := []string{"update"}

	if !slices.Contains(oneArgsCommands, command) && !slices.Contains(twoArgsCommands, command) && !slices.Contains(threeArgsCommands, command) {
		return false, errors.New("Unknown action.")
	}

	switch argsLength {
	case 0:
		return false, errors.New("This command requires at least one argument.")
	case 1:
		if !slices.Contains(oneArgsCommands, command) {
			return false, errors.New("Arguments missing.")
		} else {
			return true, nil
		}
	case 2:
		if !slices.Contains(twoArgsCommands, command) {
			return false, errors.New("Arguments missing.")
		} else {
			if command == "add" {
				return true, nil
			} else if command == "list" {
				if isTaskState(args[1]) {
					return true, nil
				} else {
					return false, errors.New("Second argument is not a valid task state")
				}
			} else if isNumber(args[1]) {
				return true, nil
			} else {
				return false, errors.New("Second argument is not an ID")
			}
		}
	case 3:
		if !slices.Contains(threeArgsCommands, command) {
			return false, errors.New("Arguments missing.")
		} else {
			if isNumber(args[1]) {
				return true, nil
			} else {
				return false, errors.New("Second argument is not an ID")
			}
		}
	default:
		return false, errors.New("Wrong number of arguments")
	}

	return true, nil
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