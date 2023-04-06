package repository

import (
	"context"

	"app/internal/model"
	"app/internal/repository/postgre"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(context.Context, model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(context.Context, model.User) error
	Delete(ctx context.Context, id int) error
}

type IBookReoisitory interface{}

type Manager struct {
	User IUserRepository
}

func NewRepository(db *gorm.DB) *Manager {
	return &Manager{
		User: postgre.NewUserRepository(db),
	}
}
