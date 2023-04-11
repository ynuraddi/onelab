package service

import (
	"context"

	"app/config"
	"app/internal/model"
	"app/internal/repository"
)

type IUserService interface {
	Create(context.Context, model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	Update(context.Context, model.User) error
	Delete(context.Context, int) error
	Auth(context.Context, model.User) error
}

type IBookService interface {
	Create(context.Context, model.Book) error
	Get(ctx context.Context, id int) (model.Book, error)
}

type IBookBorrowHistory interface {
	BorrowBook(context.Context, model.BookBorrowHistory) error
	ReturnBook(context.Context, model.BookBorrowHistory) error
	Get(ctx context.Context, id int) (model.BookBorrowHistory, error)
	ListDebtors(context.Context) ([]*model.Debtor, error)
	BookRentalForMonth(ctx context.Context, month, year int) ([]*model.UserRentalBooks, error)
}

type IAuthService interface {
	GenerateToken(username string) (string, error)
}

type Service struct {
	JWT        IAuthService
	User       IUserService
	Book       IBookService
	BookBorrow IBookBorrowHistory
}

func NewService(repo *repository.Manager, conf *config.Config) *Service {
	user := NewUserService(repo.User)

	return &Service{
		JWT:        NewJWT(conf, user),
		User:       user,
		Book:       NewBookService(repo.Book),
		BookBorrow: NewBookBorrowHistoryService(repo.BookBorrowHistory),
	}
}
