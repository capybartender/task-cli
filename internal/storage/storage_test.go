package storage

import (
	"testing"
	"github.com/capybartender/task-cli/internal/test_helpers"
)

func TestInit(t *testing.T) {
	t.Run("file path is set from function parameter", func(t *testing.T) {
		got := Init("test.json").filePath
		want := "test.json"
	 	testhelpers.Assert(t, got, want)
	})
	t.Run("file path is default", func(t *testing.T) {
		got := Init("").filePath
		want := "tasks.json"
	 	testhelpers.Assert(t, got, want)
	})
}