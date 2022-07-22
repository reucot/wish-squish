package service

import (
	"context"
	"github.com/reucot/wish-squish/internal/storage/repository"
)

type WishService struct {
	repo repository.WishEntity
}

func NewWishEntityService(repo repository.WishEntity) *WishService {
	return &WishService{
		repo: repo,
	}
}

func (w WishService) Create(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
