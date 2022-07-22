package dto

import "github.com/reucot/wish-squish/internal/models"

type SignUp struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
}

func (su SignUp) Model() *models.User {
	return &models.User{}
}

type SignIn struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
}

func (su SignIn) Model() *models.User {
	return &models.User{}
}
