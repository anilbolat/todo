package postgres

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	"github.com/heppu/todo"
	_ "github.com/lib/pq"
)

type List struct {
	dbMap *gorp.DbMap
}

func NewList() (*List, error) {
	db, err := sql.Open("postgres", "user=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(todo.Task{}, "tasks").SetKeys(true, "ID")

	if err = dbMap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}

	return &List{
		dbMap: dbMap,
	}, nil
}

func (l *List) Add(data todo.TaskData) (id uint64, err error) {
	if err = todo.ValidateTaskData(data); err != nil {
		return
	}

	task := &todo.Task{TaskData: data}
	err = l.dbMap.Insert(task)

	return task.ID, err
}

func (l *List) TaskByID(id uint64) (todo.Task, error) {
	task, err := l.dbMap.Get(todo.Task{}, id)

	// Query failed for some reason
	if err != nil {
		return todo.Task{}, err
	}

	// Query was successful but Task with given ID didn't exist
	if task == nil {
		return todo.Task{}, todo.ErrTaskNotFound
	}

	return *task.(*todo.Task), nil
}
