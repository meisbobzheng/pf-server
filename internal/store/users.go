package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sz-whereable/pants/internal/types"
)

type UserStore struct {
	db *sqlx.DB
}

func (s *UserStore) Create(ctx context.Context, user *types.User) error {
	query := `
		INSERT INTO users (username, email, password) 
    VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowxContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
