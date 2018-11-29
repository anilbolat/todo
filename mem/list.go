package mem

import (
	"sync"
	"sync/atomic"

	"github.com/heppu/todo"
)

type List struct {
	prevID uint64
	tasks  *sync.Map
}

func NewList() *List {
	return &List{
		prevID: 0,
		tasks:  &sync.Map{},
	}
}

func (tl *List) Add(data todo.TaskData) (id uint64, err error) {
	if err = todo.ValidateTaskData(data); err != nil {
		return
	}

	id = atomic.AddUint64(&tl.prevID, 1)
	tl.tasks.Store(id, data)
	return
}

func (tl *List) TaskByID(id uint64) (todo.Task, error) {
	taskData, ok := tl.tasks.Load(id)
	if !ok {
		return todo.Task{}, todo.ErrTaskNotFound
	}

	return todo.Task{
		ID:       id,
		TaskData: taskData.(todo.TaskData),
	}, nil
}
