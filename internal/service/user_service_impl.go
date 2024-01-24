package service

import (
	"context"

	"github.com/Zainal21/go-bone/internal/entity"
	"github.com/Zainal21/go-bone/internal/repositories"
)

type userServiceImpl struct {
	repo repositories.UserRepository
}

func (u userServiceImpl) ListUser(ctx context.Context) (*[]entity.User, error) {
	return u.repo.ListUser(ctx)
}

func NewUserServiceImpl(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
