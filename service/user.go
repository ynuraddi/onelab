package service

import (
	"context"
	"fmt"

	"app/model"
	"app/repository"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

const userServicePath = `userService: %w`

func (s *userService) Authenticate(ctx context.Context, user model.LogInRq) error {
	dbuser, err := s.GetByLogin(ctx, user.Login)
	if err != nil {
		return fmt.Errorf(userServicePath, err)
	}

	if err := s.comparePassword(user.Password, dbuser.Password); err != nil {
		return fmt.Errorf(userServicePath, model.ErrUserWrongPassword)
	}

	return nil
}

func (s *userService) Create(ctx context.Context, user model.CreateUserRq) error {
	hashPass, err := s.hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf(userServicePath, err)
	}

	user.Password = hashPass

	if err := s.repo.Create(ctx, user); err != nil {
		return fmt.Errorf(userServicePath, err)
	}

	return nil
}

func (s *userService) Get(ctx context.Context, id int) (model.User, error) {
	u, err := s.repo.Get(ctx, id)
	if err != nil {
		return u, fmt.Errorf(userServicePath, err)
	}

	return u, nil
}

// ADD activate User
func (s *userService) Update(ctx context.Context, user model.UpdateUserRq) error {
	dbuser, err := s.Get(ctx, user.ID)
	if err != nil {
		return fmt.Errorf(userServicePath, err)
	}

	if user.Login != "" {
		dbuser.Login = user.Login
	}
	if user.Name != "" {
		dbuser.Name = user.Name
	}

	if err := s.repo.Update(ctx, dbuser); err != nil {
		return fmt.Errorf(userServicePath, err)
	}

	return nil
}

func (s *userService) Delete(ctx context.Context, id int) error {
	if _, err := s.Get(ctx, id); err != nil {
		return fmt.Errorf(userServicePath, err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf(userServicePath, err)
	}

	return nil
}

func (s *userService) GetByLogin(ctx context.Context, login string) (model.User, error) {
	u, err := s.repo.GetByLogin(ctx, login)
	if err != nil {
		return u, fmt.Errorf(userServicePath, err)
	}

	return u, nil
}

func (s *userService) IsVerified(ctx context.Context, login string) (isActive bool, err error) {
	isActive, err = s.repo.IsVerified(ctx, login)
	if err != nil {
		return isActive, fmt.Errorf(userServicePath, err)
	}

	return isActive, nil
}

func (s *userService) comparePassword(pass string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}

func (s *userService) hashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(hash), err
}
