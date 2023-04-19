package service

import (
	"context"
	"fmt"
	"time"

	"app/model"
)

type libraryService struct {
	borrS IBookBorrowService
	userS IUserService
	bookS IBookService
	tranS ITransactionService
}

func NewLibraryService(bbs IBookBorrowService, us IUserService, bs IBookService) *libraryService {
	return &libraryService{
		borrS: bbs,
		userS: us,
		bookS: bs,
	}
}

const libraryServicePath = `libraryService: %w`

func (s *libraryService) BorrowBook(ctx context.Context, record model.LibraryBorrowRq) (rp model.LibraryBorrowRp, err error) {
	book, err := s.bookS.GetByTitle(ctx, record.BookTitle)
	if err != nil {
		return rp, fmt.Errorf(libraryServicePath, err)
	}

	user, err := s.userS.GetByLogin(ctx, record.UserLogin)
	if err != nil {
		return rp, fmt.Errorf(libraryServicePath, err)
	}

	amount := record.RentTerm * book.Price

	uuid, err := s.tranS.Create(ctx, model.CreateTransactionRq{
		UserID: user.ID,
		BookID: book.ID,
		Amount: amount,
	})
	if err != nil {
		return model.LibraryBorrowRp{}, fmt.Errorf(libraryServicePath, err)
	}

	if err := s.borrS.Create(ctx, model.CreateBookBorrowRq{
		UUID:       uuid.UUID,
		BookID:     book.ID,
		UserID:     user.ID,
		BorrowDate: time.Now(),
	}); err != nil {
		if err = s.tranS.Rollback(ctx, model.RollbackTransactionRq{UUID: uuid.UUID}); err != nil {
			return rp, fmt.Errorf(libraryServicePath, err)
		}

		return rp, fmt.Errorf(libraryServicePath, err)
	}

	return rp, nil
}
func (s *libraryService) ReturnBook(ctx context.Context) {}

func (s *libraryService) ListDebtors(ctx context.Context) (debtors []*model.LibraryDebtor, err error) {
	debtors, err = s.borrS.GetDebtors(ctx)
	if err != nil {
		return debtors, fmt.Errorf(libraryServicePath, err)
	}

	for i := 0; i < len(debtors); i++ {
		user, err := s.userS.Get(ctx, debtors[i].UserID)
		if err != nil {
			return debtors, fmt.Errorf(libraryServicePath, err)
		}

		book, err := s.bookS.Get(ctx, debtors[i].UserID)
		if err != nil {
			return debtors, fmt.Errorf(libraryServicePath, err)
		}

		debtors[i].UserName = user.Name
		debtors[i].BookTitle = book.Title
	}

	return debtors, nil
}

func (s *libraryService) ListMetric(ctx context.Context, month int) (metric []*model.LibraryMetric, err error) {
	metric, err = s.borrS.GetMetric(ctx, month)
	if err != nil {
		return metric, fmt.Errorf(libraryServicePath, err)
	}

	for i := 0; i < len(metric); i++ {
		user, err := s.userS.Get(ctx, metric[i].UserID)
		if err != nil {
			return metric, fmt.Errorf(libraryServicePath, err)
		}

		metric[i].UserName = user.Name
	}

	return metric, nil
}
