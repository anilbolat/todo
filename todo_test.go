package todo_test

import (
	"strings"
	"testing"
	"time"

	"github.com/heppu/todo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	todoList := todo.NewList()
	taskData := todo.TaskData{
		Name:      "Test task",
		CreatedAt: time.Now(),
	}

	id, err := todoList.Add(taskData)
	require.NoError(t, err, "Failed to create ")

	task, err := todoList.TaskByID(id)
	require.NoError(t, err, "Failed to get task with id %d with error: %s", id, err)
	assert.Equal(t, id, task.ID)
	assert.Equal(t, taskData, task.TaskData)

	_, err = todoList.TaskByID(1)
	require.Equal(t, todo.ErrTaskNotFound, err)
}

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

func BenchmarkList(b *testing.B) {
	todoList := todo.NewList()
	id, err := todoList.Add(todo.TaskData{
		Name:      "my task",
		CreatedAt: time.Now(),
	})
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		_, err := todoList.TaskByID(id)
		if err != nil {
			b.Fatal(err)
		}
	}
}
