package models

type Photo struct {
	Id        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Path      string `db:"path" json:"path"`
	Type      string `db:"type" json:"type"`
	ModelId   string `db:"model_id" json:"model_id"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
