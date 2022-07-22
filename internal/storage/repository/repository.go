package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type WishEntity interface {
	Create(ctx context.Context) error
}

type Repository struct {
	Authorization
	WishEntity
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationPsql(db),
		WishEntity:    NewWishPsql(db),
	}
}
