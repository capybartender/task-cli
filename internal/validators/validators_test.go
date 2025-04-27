package validators

import (
	"reflect"
	"testing"

	"github.com/capybartender/task-cli/internal/models"
)

func TestIsNumber(t *testing.T) {
	t.Run("checking if \"123\" is a number", func(t *testing.T) {
		got := isNumber("123")
		want := true
		assertBoolean(t, got, want)
	})
	t.Run("checking if \"1 2 3\" is a number", func(t *testing.T) {
		got := isNumber("1 2 3")
		want := false
		assertBoolean(t, got, want)
	})
	t.Run("checking if \"abc\" is a number", func(t *testing.T) {
		got := isNumber("abc")
		want := false
		assertBoolean(t, got, want)
	})
}

func TestIsTaskState(t *testing.T) {
	t.Run("checking if \"todo\" is a valid task state", func(t *testing.T) {
		got := isTaskState("todo")
		want := true
		assertBoolean(t, got, want)
	})
	t.Run("checking if \"in-progress\" is a valid task state", func(t *testing.T) {
		got := isTaskState("in-progress")
		want := true
		assertBoolean(t, got, want)
	})
	t.Run("checking if \"done\" is a valid task state", func(t *testing.T) {
		got := isTaskState("done")
		want := true
		assertBoolean(t, got, want)
	})
	t.Run("checking if \"invalid\" is a valid task state", func(t *testing.T) {
		got := isTaskState("invalid")
		want := false
		assertBoolean(t, got, want)
	})
}

func TestValidateArgs(t *testing.T) {
	t.Run("checking if 'add \"Buy groceries\"' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"add", "\"Buy groceries\""})
		want := &models.Command{Command: "add", Args: []string{"\"Buy groceries\""}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'update 1 \"Buy groceries and cook dinner\"' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"update", "1", "\"Buy groceries and cook dinner\""})
		want := &models.Command{Command: "update", Args: []string{"1", "\"Buy groceries and cook dinner\""}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'delete 1' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"delete", "1"})
		want := &models.Command{Command: "delete", Args: []string{"1"}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'mark-in-progress 1' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"mark-in-progress", "1"})
		want := &models.Command{Command: "mark-in-progress", Args: []string{"1"}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'mark-done 1' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"mark-done", "1"})
		want := &models.Command{Command: "mark-done", Args: []string{"1"}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list"})
		want := &models.Command{Command: "list", Args: []string{}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list done' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "done"})
		want := &models.Command{Command: "list", Args: []string{"done"}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list todo' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "todo"})
		want := &models.Command{Command: "list", Args: []string{"todo"}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list in-progress' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "in-progress"})
		want := &models.Command{Command: "list", Args: []string{"in-progress"}}
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list invalid' is an invalid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "invalid"})
		var want *models.Command = nil
		assertCorrectValidity(t, got, want)
	})
}

func assertBoolean(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func assertCorrectValidity(t testing.TB, got, want *models.Command) {
	t.Helper()
	if got == nil && want == nil {
		return		
	}
	if (got == nil && want != nil) || (got != nil && want == nil) {
		t.Errorf("got %v want %v", got, want)
		return
	}
	if got.Command != want.Command || !reflect.DeepEqual(got.Args, want.Args) {
		t.Errorf("got %v want %v", got, want)
	}
}