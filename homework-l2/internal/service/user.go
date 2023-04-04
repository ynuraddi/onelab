package service

import (
	"context"
	"fmt"

	"app/internal/model"
	"app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, user model.User) error {
	hashPass, err := s.hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("userService(Create): %w", err)
	}

	user.Password = hashPass

	if err := s.repo.Create(ctx, user); err != nil {
		return fmt.Errorf("userService(Create): %w", err)
	}

	return nil
}

func (s *userService) Get(ctx context.Context, id int) (model.User, error) {
	u, err := s.repo.Get(ctx, id)
	if err != nil {
		return u, fmt.Errorf("userService(Get): %w", err)
	}

	return u, nil
}

func (s *userService) Update(ctx context.Context, u model.User) error {
	du, err := s.repo.Get(ctx, u.ID)
	if err != nil {
		return fmt.Errorf("userService(Update): %w", err)
	}

	if u.Login != "" {
		du.Login = u.Login
	}
	if u.Name != "" {
		du.Name = u.Name
	}

	if err := s.repo.Update(ctx, du); err != nil {
		return fmt.Errorf("userService(Update): %w", err)
	}

	return nil
}

func (s *userService) Delete(ctx context.Context, u model.User) error {
	return s.repo.Delete(ctx, u)
}

func (s *userService) hashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", fmt.Errorf("userService(hashPassword): %w", err)
	}

	return string(hash), nil
}
