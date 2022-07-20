package models

type Category struct {
	Id        string `db:"id" json:"id"`
	UserId    string `db:"user_id" json:"user_id"`
	Name      string `db:"name" json:"name"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
