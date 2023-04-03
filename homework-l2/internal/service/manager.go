package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type IUserService interface {
	Create(user model.User) error
	Get(id int) (model.User, error)
}

type Service struct {
	User IUserService
}

func NewService(repo *repository.Manager) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
