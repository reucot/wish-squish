package models

type Transaction struct {
	Id         string `db:"id" json:"id"`
	WishId     string `db:"wish_id" json:"wish_id"`
	FromUserId string `db:"from_user_id" json:"from_user_id"`
	ToUserId   string `db:"to_user_id" json:"to_user_id"`
	Comment    string `db:"comment" json:"comment"`
	CreatedAt  string `db:"created_at" json:"created_at"`
}
