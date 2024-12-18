package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sz-whereable/pants/internal/types"
)

type SessionStore struct {
	db *sqlx.DB
}

func (s *SessionStore) Create(ctx context.Context, session *types.Session) error {
	query := `
		INSERT INTO sessions (user_id, token_id, token_secret) 
    VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowxContext(
		ctx,
		query,
		session.UserID,
		session.TokenID,
		session.TokenSecret,
	).Scan(
		&session.ID,
		&session.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
