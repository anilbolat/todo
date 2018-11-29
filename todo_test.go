package todo_test

import (
	"strings"
	"testing"

	"github.com/heppu/todo"
	"github.com/stretchr/testify/assert"
)

func TestValidateTask(t *testing.T) {
	tests := []struct {
		name string
		data todo.TaskData
		err  error
	}{
		{
			name: "valid name length",
			data: todo.TaskData{Name: "my task"},
			err:  nil,
		}, {
			name: "max length name",
			data: todo.TaskData{Name: strings.Repeat("a", todo.MaxTaskNameLength)},
			err:  nil,
		}, {
			name: "too long name",
			data: todo.TaskData{Name: strings.Repeat("a", todo.MaxTaskNameLength+1)},
			err:  todo.ErrTooLongTaskName,
		}, {
			name: "empty name",
			data: todo.TaskData{},
			err:  todo.ErrEmptyTaskName,
		}, {
			name: "max length utf8",
			data: todo.TaskData{Name: strings.Repeat("€", todo.MaxTaskNameLength)},
			err:  nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := todo.ValidateTaskData(test.data)
			assert.Equal(t, test.err, err)
		})
	}
}
