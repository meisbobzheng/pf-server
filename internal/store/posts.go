package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sz-whereable/pants/internal/types"
)

type PostStore struct {
	db *sqlx.DB
}

func (s *PostStore) Create(ctx context.Context, post *types.Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags) 
    VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowxContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		post.Tags,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
