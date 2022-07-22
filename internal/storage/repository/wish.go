package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type WishPsql struct {
	db *sqlx.DB
}

func NewWishPsql(db *sqlx.DB) *WishPsql {
	return &WishPsql{
		db: db,
	}
}

func (w WishPsql) Create(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
