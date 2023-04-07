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
	Create(context.Context, model.BookBorrowHistory) error
	Get(ctx context.Context, id int) (model.BookBorrowHistory, error)
	ListDebtors(context.Context) ([]*model.Debtor, error)
	BookRentalForMonth(ctx context.Context, month, year int) ([]*model.UserRentalBooks, error)
}

type Manager struct {
	User              IUserRepository
	Book              IBookReoisitory
	BookBorrowHistory IBookBorrowHistory
}

func NewRepository(db *gorm.DB) *Manager {
	return &Manager{
		User: postgre.NewUserRepository(db),
	}
}
