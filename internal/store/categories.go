package store

import (
	"context"
	"database/sql"
)

type Category struct {
	Id   int64      `json:"id"`
	Name NullString `json:"name"`
}

type CategoryStore struct {
	db *sql.DB
}

func (s *CategoryStore) Create(ctx context.Context, payload Category) error {
	query := `
		INSERT INTO categories (name)
		VALUES (?)
	`
	if _, err := s.db.ExecContext(ctx, query, payload.Name); err != nil {
		return err
	}
	return nil
}
