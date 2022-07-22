package service

import (
	"github.com/reucot/wish-squish/internal/storage/repository"
)

type AuthorizationService struct {
	ap repository.Authorization
}

func NewAuthorizationService(ap repository.Authorization) *AuthorizationService {
	return &AuthorizationService{
		ap: ap,
	}
}

//func (s AuthorizationService) SignUp(ctx context.Context, su *models.SignUp) error {
//	return nil
//}

func (s AuthorizationService) SignIn() {

}
