package service

import (
	"context"

	"app/config"
	"app/model"
	"app/repository"
)

type IUserService interface {
	Authenticate(ctx context.Context, user model.LogInRq) error

	Create(ctx context.Context, user model.CreateUserRq) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.UpdateUserRq) error
	Delete(ctx context.Context, id int) error

	GetByLogin(ctx context.Context, login string) (model.User, error)
	IsVerified(ctx context.Context, login string) (isActive bool, err error)
}

//go:generate mockery --name IBookService
type IBookService interface {
	Create(ctx context.Context, book model.CreateBookRq) error
	Get(ctx context.Context, id int) (model.Book, error)
	Update(ctx context.Context, book model.UpdateBookRq) error
	Delete(ctx context.Context, id int) error
}

type IBookBorrowHistory interface {
	Create(ctx context.Context, record model.CreateBookBorrowRq) error
	Get(ctx context.Context, id int) (model.BookBorrow, error)
	Update(ctx context.Context, record model.UpdateBookBorrowRq) error
	Delete(ctx context.Context, id int) error

	ListDebtors(ctx context.Context) (debtors []*model.BookBorrowDebtorRp, err error)
	ListMetric(ctx context.Context, month int) (metric []*model.BookBorrowMetricRp, err error)
}

type ITransactionService interface{}

type Manager struct {
	User       IUserService
	Book       IBookService
	BookBorrow IBookBorrowHistory
}

func NewService(repo *repository.Manager, conf *config.Config) *Manager {
	us := NewUserService(repo.User)
	bs := NewBookService(repo.Book)
	bbs := NewBookBorrowService(repo.BookBorrow, us, bs)

	return &Manager{
		User:       us,
		Book:       bs,
		BookBorrow: bbs,
	}
}
