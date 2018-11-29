package todo

import (
	"errors"
	"time"
	"unicode/utf8"
)

const MaxTaskNameLength = 50

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrEmptyTaskName   = errors.New("empty task name")
	ErrTooLongTaskName = errors.New("too long task name")
)

type Task struct {
	ID uint64
	TaskData
}

type TaskData struct {
	Name      string // Number of characters must be greater than 0 but less than 50
	CreatedAt time.Time
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
