package dto

import "github.com/reucot/wish-squish/internal/entity"

type SignUp struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
}

func (su SignUp) Model() *entity.SignUp {
	return &entity.SignUp{
		User: entity.User{
			Email:    su.Email,
			Password: su.Password,
		},
	}
}

type SignIn struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
}

func (su SignIn) Model() *entity.SignIn {
	return &entity.SignIn{
		User: entity.User{
			Email:    su.Email,
			Password: su.Password,
		},
	}
}
