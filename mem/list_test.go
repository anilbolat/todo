package mem_test

import (
	"testing"
	"time"

	"github.com/heppu/todo"
	"github.com/heppu/todo/mem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	todoList := mem.NewList()
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

	_, err = todoList.TaskByID(99999)
	require.Equal(t, todo.ErrTaskNotFound, err)

	_, err = todoList.Add(todo.TaskData{})
	require.Equal(t, todo.ErrEmptyTaskName, err, "Adding task with empty name didn't fail")
}

func BenchmarkList(b *testing.B) {
	todoList := mem.NewList()
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
