package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	Name string `json:"name"`
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
		SELECT name
		FROM tasks
		LIMIT 10
	`)

	if err != nil {
		panic(err.Error())
	}

	var result []Task
	row := Task{}
	for rows.Next() {
		if err := rows.Scan(
			&row.Name,
		); err != nil {
			log.Fatal(err)
		}

		result = append(result, row)
	}

	return result, nil
}
