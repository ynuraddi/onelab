package service

import "app/internal/repository"

type IUserService interface {
	Create()
	Get()
	Update()
	Delete()
}

type Service struct {
	User IUserService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
