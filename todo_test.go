package todo_test

import (
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
