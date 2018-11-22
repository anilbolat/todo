package todo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/heppu/todo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListAPI(t *testing.T) {
	handler := todo.NewHandler(todo.NewList())

	for i := 0; i < 2; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			taskData := todo.TaskData{
				Name:      "my task",
				CreatedAt: time.Now().Round(0), // strip monotonic clock
			}
			body, err := json.Marshal(taskData)
			require.NoError(t, err, "Failed to marshal task data")

			req := httptest.NewRequest("POST", "http://localhost/api/tasks", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Result().StatusCode, "Unexpected status code after posting new task")

			var id uint64
			err = json.NewDecoder(w.Result().Body).Decode(&id)
			require.NoError(t, err, "Failed to decode response body to uint64")

			req = httptest.NewRequest("GET", fmt.Sprintf("http://localhost/api/tasks/%d", id), nil)
			w = httptest.NewRecorder()

			handler.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Result().StatusCode, "Unexpected status code after getting task")

			task := todo.Task{}
			err = json.NewDecoder(w.Result().Body).Decode(&task)
			require.NoError(t, err, "Failed to decode response body to task")
			assert.Equal(t, todo.Task{ID: id, TaskData: taskData}, task, "Task data doesn't match")
		})
	}
}

func TestAPIServer(t *testing.T) {
	handler := todo.NewHandler(todo.NewList())
	server := httptest.NewServer(handler)
	defer server.Close()

	taskData := todo.TaskData{
		Name:      "my task",
		CreatedAt: time.Now().Round(0), // strip monotonic clock
	}
	body, err := json.Marshal(taskData)
	require.NoError(t, err, "Failed to marshal task data")

	resp, err := http.Post(server.URL+"/api/tasks", "", bytes.NewBuffer(body))
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code after posting new task")

	var id uint64
	err = json.NewDecoder(resp.Body).Decode(&id)
	require.NoError(t, err, "Failed to decode response body to uint64")

	resp, err = http.Get(fmt.Sprintf("%s/api/tasks/%d", server.URL, id))
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code after getting task")

	task := todo.Task{}
	err = json.NewDecoder(resp.Body).Decode(&task)
	require.NoError(t, err, "Failed to decode response body to task")
	assert.Equal(t, todo.Task{ID: id, TaskData: taskData}, task, "Task data doesn't match")
}
