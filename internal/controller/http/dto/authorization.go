package dto

import "github.com/reucot/wish-squish/internal/models"

type SignUp struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
}

func (su SignUp) Model() *models.SignUp {
	return &models.SignUp{
		User: models.User{
			Email:    su.Email,
			Password: su.Password,
		},
	}
}

type SignIn struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
}

func (su SignIn) Model() *models.SignIn {
	return &models.SignIn{
		User: models.User{
			Email:    su.Email,
			Password: su.Password,
		},
	}
}
