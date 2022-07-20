package models

type WishCategory struct {
	Id         string `db:"id" json:"id"`
	CategoryId string `db:"category_id" json:"category_id"`
	WishId     string `db:"wish_id" json:"wish_id"`
}
