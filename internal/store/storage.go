package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tasks interface {
		GetAll(context.Context) ([]Task, error)
		GetByID(context.Context, int64) (Task, error)
	}
	Categories interface {
		Create(context.Context, Category) (error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TaskStore{db},
		Categories: &CategoryStore{db},
	}
}
