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

	GetByTitle(ctx context.Context, title string) (model.Book, error)
}

type IBookBorrowService interface {
	Create(ctx context.Context, record model.CreateBookBorrowRq) error
	Get(ctx context.Context, id int) (model.BookBorrow, error)
	Update(ctx context.Context, record model.UpdateBookBorrowRq) error
	Delete(ctx context.Context, id int) error

	GetDebtors(ctx context.Context) ([]*model.LibraryDebtor, error)
	GetMetric(ctx context.Context, month int) ([]*model.LibraryMetric, error)
}

type ILibraryService interface {
	BorrowBook(ctx context.Context, record model.LibraryBorrowRq) (model.LibraryBorrowRp, error)
	ReturnBook(ctx context.Context)

	ListDebtors(ctx context.Context) (debtors []*model.LibraryDebtor, err error)
	ListMetric(ctx context.Context, month int) (metric []*model.LibraryMetric, err error)
}

type ITransactionService interface {
	Create(ctx context.Context, tr model.CreateTransactionRq) (model.CreateTransactionRp, error)
	Pay(ctx context.Context, tr model.PayTransactionRq) error
}

type Manager struct {
	User       IUserService
	Book       IBookService
	BookBorrow IBookBorrowService
}

func NewService(conf *config.Config, repo *repository.Manager) *Manager {
	userS := NewUserService(repo.User)
	bookS := NewBookService(repo.Book)
	borrS := NewBookBorrowService(repo.BookBorrow, userS, bookS)

	return &Manager{
		User:       userS,
		Book:       bookS,
		BookBorrow: borrS,
	}
}
