package service

import (
	"fmt"

	"app/internal/model"
	"app/internal/repository"
)

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(user model.User) error {
	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("userService(Create): %w", err)
	}

	return nil
}

func (s *userService) Get(id int) (model.User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return user, fmt.Errorf("userService(Get): %w", err)
	}

	return user, nil
}
