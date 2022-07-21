package models

type UserSettings struct {
	Id           string `db:"id" json:"id"`
	UserId       string `db:"user_id" json:"user_id"`
	Handle       string `db:"handle" json:"handle"`
	WishListName string `db:"wish_list_name" json:"wish_list_name"`
	Alias        string `db:"alias" json:"alias"`
}
