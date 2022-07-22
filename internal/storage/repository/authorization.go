package repository

import "github.com/jmoiron/sqlx"

type AuthorizationPsql struct {
	db *sqlx.DB
}

func NewAuthorizationPsql(db *sqlx.DB) *AuthorizationPsql {
	return &AuthorizationPsql{db: db}
}
