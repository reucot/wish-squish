package models

type User struct {
	Id         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Surname    string `db:"surname" json:"surname"`
	Email      string `db:"email" json:"email"`
	Password   string `json:"password"`
	CreatedAt  string `db:"created_at" json:"created_at"`
	UpdatedAt  string `db:"updated_at" json:"updated_at"`
	VerifiedAt string `db:"verified_at" json:"verified_at"`
	RoleId     string `db:"role_id" json:"role_id"`
}
