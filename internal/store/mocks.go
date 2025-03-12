package store

import (
	"context"
	"database/sql"
)

func NewMockStorage() Storage {
	return Storage{
		Tasks: &MockTaskStore{},
	}
}

type MockTaskStore struct{}

func (s *MockTaskStore) GetAll(ctx context.Context) ([]Task, error) {
	return []Task{{1, NullString{sql.NullString{String: "name", Valid: true}}}}, nil
}

func (s *MockTaskStore) GetByID(ctx context.Context, userID int64) (Task, error) {
	return Task{}, nil
}
