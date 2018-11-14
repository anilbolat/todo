package todo

import (
	"errors"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type List struct {
	nextID uint64
	tasks  map[uint64]TaskData
}

type Task struct {
	ID uint64
	TaskData
}

type TaskData struct {
	Name      string
	CreatedAt time.Time
}

func NewList() *List {
	return &List{
		nextID: 0,
		tasks:  make(map[uint64]TaskData),
	}
}

func (tl *List) Add(data TaskData) (id uint64, err error) {
	id = tl.nextID
	tl.tasks[id] = data
	tl.nextID++
	return
}

func (tl *List) TaskByID(id uint64) (Task, error) {
	taskData, ok := tl.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return Task{
		ID:       id,
		TaskData: taskData,
	}, nil
}
