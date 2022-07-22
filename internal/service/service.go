package service

import (
	"context"
	"github.com/reucot/wish-squish/internal/storage/repository"
)

type Authorization interface {
}

type WishEntity interface {
	Create(ctx context.Context) error
}

type Service struct {
	Authorization
	WishEntity
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.Authorization),
		WishEntity:    NewWishEntityService(repos.WishEntity),
	}
}
