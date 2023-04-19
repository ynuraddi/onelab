package service

import (
	"context"
	"fmt"

	"app/model"
	"app/repository"
)

type bookService struct {
	repo repository.IBookRepository
}

func NewBookService(repo repository.IBookRepository) *bookService {
	return &bookService{
		repo: repo,
	}
}

const bookServicePath = `bookService: %w`

func (s *bookService) Create(ctx context.Context, b model.CreateBookRq) error {
	if err := s.repo.Create(ctx, b); err != nil {
		return fmt.Errorf(bookServicePath, err)
	}
	return nil
}

func (s *bookService) Get(ctx context.Context, id int) (model.Book, error) {
	book, err := s.repo.Get(ctx, id)
	if err != nil {
		return book, fmt.Errorf(bookServicePath, err)
	}
	return book, nil
}

func (s *bookService) Update(ctx context.Context, book model.UpdateBookRq) error {
	dbbook, err := s.Get(ctx, book.ID)
	if err != nil {
		return fmt.Errorf(bookServicePath, err)
	}

	if book.Author != "" {
		dbbook.Author = book.Author
	}
	if book.Title != "" {
		dbbook.Title = book.Title
	}

	if err := s.repo.Update(ctx, dbbook); err != nil {
		return fmt.Errorf(bookServicePath, err)
	}
	return nil
}

func (s *bookService) Delete(ctx context.Context, id int) error {
	if _, err := s.Get(ctx, id); err != nil {
		return fmt.Errorf(bookServicePath, err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf(bookServicePath, err)
	}
	return nil
}

func (s *bookService) GetByTitle(ctx context.Context, title string) (model.Book, error) {
	u, err := s.repo.GetByTitle(ctx, title)
	if err != nil {
		return u, fmt.Errorf(bookServicePath, err)
	}

	return u, nil
}
