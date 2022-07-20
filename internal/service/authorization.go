package service

import (
	"github.com/reucot/wish-squish/internal/storage/psql"
)

type AuthorizationService struct {
	ap psql.AuthorizationPsql
}

func NewAuthorizationService(ap *psql.AuthorizationPsql) *AuthorizationService {
	return &AuthorizationService{
		ap: *ap,
	}
}

//func (s AuthorizationService) SignUp(ctx context.Context, su *models.SignUp) error {
//	return nil
//}

func (s AuthorizationService) SignIn() {

}
