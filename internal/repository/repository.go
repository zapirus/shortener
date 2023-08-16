package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	InsertURL(beforeURL, afterURL string) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (*repository, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return &repository{
		db: db,
	}, nil
}

func (r *repository) InsertURL(beforeURL, afterURL string) error {
	query := "INSERT INTO urls (before, after) VALUES (?, ?)"
	_, err := r.db.Exec(query, beforeURL, afterURL)
	if err != nil {
		return fmt.Errorf("failed to insert URL: %w", err)
	}
	return nil
}
