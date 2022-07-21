package psql

import "github.com/jmoiron/sqlx"

type AuthorizationPsql struct {
	db *sqlx.DB
}

func NewAuthorizationPsql() *AuthorizationPsql {
	return &AuthorizationPsql{}
}
