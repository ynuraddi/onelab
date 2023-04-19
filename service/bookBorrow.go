package service

import (
	"context"
	"fmt"
	"time"

	"app/model"
	"app/repository"
)

type bookBorrowService struct {
	repo  repository.IBookBorrowRepository
	userS IUserService
	bookS IBookService
}

func NewBookBorrowService(repo repository.IBookBorrowRepository, us IUserService, bs IBookService) *bookBorrowService {
	return &bookBorrowService{
		repo:  repo,
		userS: us,
		bookS: bs,
	}
}

const bookBorrowServicePath = `bookBorrowService: %w`

var nilTime time.Time

func (s *bookBorrowService) Create(ctx context.Context, record model.CreateBookBorrowRq) error {
	if _, err := s.userS.Get(ctx, record.UserID); err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}

	if _, err := s.bookS.Get(ctx, record.BookID); err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}

	if record.BorrowDate == nilTime {
		record.BorrowDate = time.Now()
	}

	if err := s.repo.Create(ctx, record); err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}
	return nil
}

func (s *bookBorrowService) Get(ctx context.Context, id int) (model.BookBorrow, error) {
	record, err := s.repo.Get(ctx, id)
	if err != nil {
		return record, fmt.Errorf(bookBorrowServicePath, err)
	}
	return record, nil
}

func (s *bookBorrowService) Update(ctx context.Context, record model.UpdateBookBorrowRq) error {
	dbrecord, err := s.Get(ctx, record.ID)
	if err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}

	nilTime := time.Time{}

	if record.BookID > 0 {
		dbrecord.BookID = record.BookID
	}
	if record.UserID > 0 {
		dbrecord.UserID = record.UserID
	}
	if record.BorrowDate != nilTime {
		dbrecord.BorrowDate = record.BorrowDate
	}
	if record.ReturnDate != nilTime {
		dbrecord.ReturnDate = record.ReturnDate
	}

	if err := s.repo.Update(ctx, dbrecord); err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}
	return nil
}

func (s *bookBorrowService) Delete(ctx context.Context, id int) error {
	if _, err := s.Get(ctx, id); err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}
	return nil
}

func (s *bookBorrowService) GetDebtors(ctx context.Context) (debtors []*model.LibraryDebtor, err error) {
	debtors, err = s.repo.GetDebtors(ctx)
	if err != nil {
		return debtors, fmt.Errorf(bookBorrowServicePath, err)
	}

	return debtors, nil
}

func (s *bookBorrowService) GetMetric(ctx context.Context, month int) (metric []*model.LibraryMetric, err error) {
	metric, err = s.repo.GetMetric(ctx, month)
	if err != nil {
		return metric, fmt.Errorf(bookBorrowServicePath, err)
	}

	return metric, nil
}

// func (s *bookBorrowService) ListDebtors(ctx context.Context) (debtors []*model.BookBorrowDebtorRp, err error) {
// 	list, err := s.repo.ListDebtors(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf(bookBorrowServicePath, err)
// 	}
// 	return list, nil
// }

// func (s *bookBorrowService) ListMetric(ctx context.Context, month int) (metric []*model.BookBorrowMetricRp, err error) {
// 	list, err := s.repo.ListMetric(ctx, month)
// 	if err != nil {
// 		return nil, fmt.Errorf(bookBorrowServicePath, err)
// 	}
// 	return list, nil
// }
