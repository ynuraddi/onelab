package service

import (
	"context"

	"app/internal/model"
	"app/internal/repository"
)

type IUserService interface {
	Create(context.Context, model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(context.Context, model.User) error
	Delete(context.Context, model.User) error
}

type Service struct {
	User IUserService
}

func NewService(repo *repository.Manager) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
