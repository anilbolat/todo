package todo

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"
)

const MaxTaskNameLength = 50

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrEmptyTaskName   = errors.New("empty task name")
	ErrTooLongTaskName = errors.New("too long task name")
)

type List struct {
	prevID uint64
	tasks  *sync.Map
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
		prevID: 0,
		tasks:  &sync.Map{},
	}
}

func (tl *List) Add(data TaskData) (id uint64, err error) {
	if err = ValidateTaskData(data); err != nil {
		return
	}

	id = atomic.AddUint64(&tl.prevID, 1)
	tl.tasks.Store(id, data)
	return
}

func (tl *List) TaskByID(id uint64) (Task, error) {
	taskData, ok := tl.tasks.Load(id)
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return Task{
		ID:       id,
		TaskData: taskData.(TaskData),
	}, nil
}

func ValidateTaskData(t TaskData) error {
	if t.Name == "" {
		return ErrEmptyTaskName
	}

	if utf8.RuneCountInString(t.Name) > MaxTaskNameLength {
		return ErrTooLongTaskName
	}

	return nil
}
