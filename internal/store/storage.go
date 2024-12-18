package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sz-whereable/pants/internal/types"
)

type Storage struct {
	Users interface {
		Create(context.Context, *types.User) error
	}
	Sessions interface {
		Create(context.Context, *types.Session) error
	}
}

func NewStorage(db *sqlx.DB) Storage {
	return Storage{
		Users:    &UserStore{db},
		Sessions: &SessionStore{db},
	}
}
