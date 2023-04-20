package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"app/model"
	"app/repository"
)

type bookBorrowService struct {
	repo  repository.IBookBorrowRepository
	bookS IBookService
}

func NewBookBorrowService(repo repository.IBookBorrowRepository, bs IBookService) *bookBorrowService {
	return &bookBorrowService{
		repo:  repo,
		bookS: bs,
	}
}

const timeLayout = "2006-01-02"

const bookBorrowServicePath = `bookBorrowService: %w`

func (s *bookBorrowService) Create(ctx context.Context, record model.CreateBookBorrowRq) error {
	borrowDate, err := time.Parse(timeLayout, record.BorrowDate)
	if err != nil {
		return fmt.Errorf(bookBorrowServicePath, err)
	}

	if err := s.repo.Create(ctx, model.CreateBookBorrowRepo{
		UUID:       record.UUID,
		BookID:     record.BookID,
		UserID:     record.UserID,
		BorrowDate: borrowDate,
	}); err != nil {
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

func (s *bookBorrowService) ListDebtors(ctx context.Context) (debtors []*model.LibraryDebtor, err error) {
	debtors, err = s.repo.ListDebtors(ctx)
	if err != nil {
		return debtors, fmt.Errorf(bookBorrowServicePath, err)
	}

	return debtors, nil
}

func (s *bookBorrowService) ListMetric(ctx context.Context, month int) (metric []*model.LibraryMetricUserBook, err error) {
	list, err := s.repo.ListMetric(ctx, month)
	if err != nil {
		return metric, fmt.Errorf(bookBorrowServicePath, err)
	}

	books := map[int]*model.Book{}

	// LibraryMetricRepo ==> LibraryMetric
	str := ""
	for _, l := range list {
		str = l.Books
		str = strings.TrimPrefix(str, "{")
		str = strings.TrimSuffix(str, "}")
		booksStrArr := strings.Split(str, ",")

		booksTitles := []string{}

		for _, strID := range booksStrArr {
			bookID, _ := strconv.Atoi(strID)
			if book, ok := books[bookID]; ok {
				booksTitles = append(booksTitles, book.Title)
				continue
			}

			book, err := s.bookS.Get(ctx, bookID)
			if err != nil {
				return metric, fmt.Errorf(bookBorrowServicePath, err)
			}
			booksTitles = append(booksTitles, book.Title)
			books[bookID] = &book
		}

		metric = append(metric, &model.LibraryMetricUserBook{
			UserID:   l.UserID,
			UserName: l.UserName,
			Books:    booksTitles,
		})
	}

	return metric, nil
}

func (s *bookBorrowService) GetByUserBook(ctx context.Context, userID, bookID int) (model.BookBorrow, error) {
	record, err := s.repo.GetByUserBook(ctx, userID, bookID)
	if err != nil {
		return record, fmt.Errorf(bookBorrowServicePath, err)
	}

	return record, nil
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
