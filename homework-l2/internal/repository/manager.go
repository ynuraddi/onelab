package repository

import "app/internal/model"

type IUserRepository interface {
	Create(model.User) error
	Get(id int) (model.User, error)
}

type Manager struct {
	User IUserRepository
}

func NewRepository() *Manager {
	return &Manager{
		User: NewUserRepository(),
	}
}
