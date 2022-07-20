package models

type Wish struct {
	Id            string `db:"id" json:"id"`
	UserId        string `db:"user_id" json:"user_id"`
	Name          string `db:"name" json:"name"`
	Price         string `db:"price" json:"price"`
	IsAllowRepeat string `db:"is_allow_repeat" json:"is_allow_repeat"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UpdatedAt     string `db:"updated_at" json:"updated_at"`
}
