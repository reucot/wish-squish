package entity

type SignUp struct {
	User User `json:"user"`
}

type SignIn struct {
	User User `json:"user"`
}
