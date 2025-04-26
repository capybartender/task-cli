package validators

import "testing"

func TestIsNumber(t *testing.T) {
	t.Run("checking if \"123\" is a number", func(t *testing.T) {
		got := isNumber("123")
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if \"1 2 3\" is a number", func(t *testing.T) {
		got := isNumber("1 2 3")
		want := false
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if \"abc\" is a number", func(t *testing.T) {
		got := isNumber("abc")
		want := false
		assertCorrectValidity(t, got, want)
	})
}

func TestIsTaskState(t *testing.T) {
	t.Run("checking if \"todo\" is a valid task state", func(t *testing.T) {
		got := isTaskState("todo")
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if \"in-progress\" is a valid task state", func(t *testing.T) {
		got := isTaskState("in-progress")
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if \"done\" is a valid task state", func(t *testing.T) {
		got := isTaskState("done")
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if \"invalid\" is a valid task state", func(t *testing.T) {
		got := isTaskState("invalid")
		want := false
		assertCorrectValidity(t, got, want)
	})
}

func TestValidateArgs(t *testing.T) {
	/*
	   # Adding a new task
	   add "Buy groceries"

	   # Updating and deleting tasks
	   update 1 "Buy groceries and cook dinner"
	   delete 1

	   # Marking a task as in progress or done
	   mark-in-progress 1
	   mark-done 1

	   # Listing all tasks
	   list

	   # Listing tasks by status
	   list done
	   list todo
	   list in-progress
	*/
	t.Run("checking if 'add \"Buy groceries\"' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"add", "\"Buy groceries\""})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'update 1 \"Buy groceries and cook dinner\"' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"update", "1", "\"Buy groceries and cook dinner\""})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'delete 1' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"delete", "1"})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'mark-in-progress 1' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"mark-in-progress", "1"})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'mark-done 1' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"mark-done", "1"})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list"})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list done' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "done"})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list todo' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "todo"})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list in-progress' is a valid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "in-progress"})
		want := true
		assertCorrectValidity(t, got, want)
	})
	t.Run("checking if 'list invalid' is an invalid command", func(t *testing.T) {
		got, _ := ValidateArgs([]string{"list", "invalid"})
		want := false
		assertCorrectValidity(t, got, want)
	})
}

func assertCorrectValidity(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}