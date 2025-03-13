package store

import (
	"context"
	"database/sql"
	"fmt"
)

type Task struct {
	Id   int64      `json:"id"`
	Name NullString `json:"name"`
}

type TaskStore struct {
	db *sql.DB
}

func (s *TaskStore) GetByID(ctx context.Context, userID int64) (Task, error) {
	return Task{}, nil
}

func (s *TaskStore) GetAll(ctx context.Context) ([]Task, error) {
	fmt.Println(ctx)
	rows, err := s.db.Query(`
		SELECT id, name
		FROM tasks
		LIMIT 10
	`)

	if err != nil {
		return nil, err
	}

	var result []Task
	row := Task{}
	for rows.Next() {
		if err := rows.Scan(&row.Id, &row.Name); err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}
