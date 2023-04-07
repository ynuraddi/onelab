package service

import (
	"context"

	"app/internal/model"
	"app/internal/repository"
)

type bookService struct {
	repo repository.IBookRepository
}

func NewBookService(repo repository.IBookRepository) *bookService {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) Create(ctx context.Context, b model.Book) error {
	return s.repo.Create(ctx, b)
}

func (s *bookService) Get(ctx context.Context, id int) (model.Book, error) {
	return s.repo.Get(ctx, id)
}
