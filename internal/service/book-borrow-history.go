package service

import (
	"context"

	"app/internal/model"
	"app/internal/repository"
)

type bookBorrowHistoryService struct {
	repo repository.IBookBorrowHistoryRepository
}

func NewBookBorrowHistoryService(repo repository.IBookBorrowHistoryRepository) *bookBorrowHistoryService {
	return &bookBorrowHistoryService{
		repo: repo,
	}
}

func (s *bookBorrowHistoryService) Create(ctx context.Context, b model.BookBorrowHistory) error {
	return s.repo.Create(ctx, b)
}

func (s *bookBorrowHistoryService) Get(ctx context.Context, id int) (model.BookBorrowHistory, error) {
	return s.repo.Get(ctx, id)
}

func (s *bookBorrowHistoryService) ListDebtors(ctx context.Context) ([]*model.Debtor, error) {
	return s.repo.ListDebtors(ctx)
}

func (s *bookBorrowHistoryService) BookRentalForMonth(ctx context.Context, month, year int) ([]*model.UserRentalBooks, error) {
	return s.repo.BookRentalForMonth(ctx, month, year)
}
