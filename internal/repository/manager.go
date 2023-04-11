package repository

import (
	"context"

	"app/config"
	"app/internal/model"
	"app/internal/repository/postgre"
)

type IUserRepository interface {
	Create(context.Context, model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	GetByUsername(context.Context, string) (model.User, error)
	Update(context.Context, model.User) error
	Delete(ctx context.Context, id int) error
}

type IBookRepository interface {
	Create(context.Context, model.Book) error
	Get(ctx context.Context, id int) (model.Book, error)
}

type IBookBorrowHistoryRepository interface {
	BorrowBook(context.Context, model.BookBorrowHistory) error
	ReturnBook(context.Context, model.BookBorrowHistory) error
	Get(ctx context.Context, id int) (model.BookBorrowHistory, error)
	ListDebtors(context.Context) ([]*model.Debtor, error)
	BookRentalForMonth(ctx context.Context, month, year int) ([]*model.UserRentalBooks, error)
}

type Manager struct {
	User              IUserRepository
	Book              IBookRepository
	BookBorrowHistory IBookBorrowHistoryRepository
}

func NewRepository(conf *config.Config) *Manager {
	db := postgre.OpenDB(conf)

	return &Manager{
		User:              postgre.NewUserRepository(db),
		Book:              postgre.NewBookRepository(db),
		BookBorrowHistory: postgre.NewBookBorrowHistoryRepository(db),
	}
}
