package service

import (
	"context"

	"github.com/S4mkiel/validation/domain/entity"
	"github.com/S4mkiel/validation/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	u, e := s.repo.Create(ctx, user)
	if e != nil {
		return nil, e
	} else {
		return u, nil
	}
}
