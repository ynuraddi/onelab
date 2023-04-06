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

// TODO
type IBookReoisitory interface {
	Create(context.Context, model.Book) error
	Get(ctx context.Context, id int) (model.Book, error)
}

// TODO
type IBookBorrowHistory interface {
	Create(context.Context)
	Get(context.Context)
	ListDebtors(context.Context)
	UsersBookCountLastMonth(context.Context)
}

type Manager struct {
	User IUserRepository
}

func NewRepository(db *gorm.DB) *Manager {
	return &Manager{
		User: postgre.NewUserRepository(db),
	}
}
