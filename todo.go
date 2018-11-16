package todo

import (
	"errors"
	"time"
)

const MaxTaskNameLength = 50

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrEmptyTaskName   = errors.New("empty task name")
	ErrTooLongTaskName = errors.New("too long task name")
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
	Name      string // Number of characters must be greater than 0 but less than 50
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

func ValidateTaskData(t TaskData) error {
	if t.Name == "" {
		return ErrEmptyTaskName
	}

	if len(t.Name) > MaxTaskNameLength {
		return ErrTooLongTaskName
	}

	return nil
}
