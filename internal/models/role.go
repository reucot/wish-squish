package models

type Role struct {
	Id   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
