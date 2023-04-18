package repository

import (
	"context"

	"app/config"
	"app/model"
	"app/repository/postgres"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.CreateUserRq) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id int) error

	GetByLogin(ctx context.Context, login string) (model.User, error)
	IsVerified(ctx context.Context, login string) (isActive bool, err error)
}

//go:generate mockery --name IBookReposiotry
type IBookRepository interface {
	Create(ctx context.Context, book model.CreateBookRq) error
	Get(ctx context.Context, id int) (model.Book, error)
	Update(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, id int) error
}

type IBookBorrowRepository interface {
	Create(ctx context.Context, record model.CreateBookBorrowRq) error
	Get(ctx context.Context, id int) (model.BookBorrow, error)
	Update(ctx context.Context, record model.BookBorrow) error
	Delete(ctx context.Context, id int) error

	ListDebtors(ctx context.Context) (debtors []*model.BookBorrowDebtorRp, err error)
	ListMetric(ctx context.Context, month int) (metric []*model.BookBorrowMetricRp, err error)
}

type Manager struct {
	User       IUserRepository
	Book       IBookRepository
	BookBorrow IBookBorrowRepository
}

func NewRepository(conf *config.Config) *Manager {
	db := postgres.OpenDB(conf)

	return &Manager{
		User:       postgres.NewUserRepository(db),
		Book:       postgres.NewBookRepository(db),
		BookBorrow: postgres.NewBookBorrowRepository(db),
	}
}
