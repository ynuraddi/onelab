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

func NewLibraryService(bbs IBookBorrowService, us IUserService, bs IBookService, trs ITransactionService) *libraryService {
	return &libraryService{
		borrS: bbs,
		userS: us,
		bookS: bs,
		tranS: trs,
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

	uuid, err := s.tranS.Create(ctx, model.CreateTransactionRq{
		UserID: user.ID,
		BookID: book.ID,
		Amount: book.Price,
	})
	if err != nil {
		return model.LibraryBorrowRp{}, fmt.Errorf(libraryServicePath, err)
	}
	rp.TransactionUUID = uuid.UUID
	rp.Score = uuid.Amount

	if err := s.borrS.Create(ctx, model.CreateBookBorrowRq{
		UUID:       uuid.UUID,
		BookID:     book.ID,
		UserID:     user.ID,
		BorrowDate: record.BorrowDate,
	}); err != nil {
		if err = s.tranS.Rollback(ctx, model.RollbackTransactionRq{UUID: uuid.UUID}); err != nil {
			return rp, fmt.Errorf(libraryServicePath, err)
		}

		return rp, fmt.Errorf(libraryServicePath, err)
	}

	return rp, nil
}

func (s *libraryService) ReturnBook(ctx context.Context, record model.LibraryReturnRq) error {
	book, err := s.bookS.GetByTitle(ctx, record.BookTitle)
	if err != nil {
		return fmt.Errorf(libraryServicePath, err)
	}

	user, err := s.userS.GetByLogin(ctx, record.UserLogin)
	if err != nil {
		return fmt.Errorf(libraryServicePath, err)
	}

	borrowRecord, err := s.borrS.GetByUserBook(ctx, user.ID, book.ID)
	if err != nil {
		return fmt.Errorf(libraryServicePath, err)
	}

	if err := s.borrS.Update(ctx, model.UpdateBookBorrowRq{
		ID:         borrowRecord.ID,
		ReturnDate: time.Now(),
	}); err != nil {
		return fmt.Errorf(libraryServicePath, err)
	}

	return nil
}

func (s *libraryService) ListDebtors(ctx context.Context) (debtors []*model.LibraryDebtor, err error) {
	debtors, err = s.borrS.ListDebtors(ctx)
	if err != nil {
		return debtors, fmt.Errorf(libraryServicePath, err)
	}

	for i := 0; i < len(debtors); i++ {
		user, err := s.userS.Get(ctx, debtors[i].UserID)
		if err != nil {
			return debtors, fmt.Errorf(libraryServicePath, err)
		}

		book, err := s.bookS.Get(ctx, debtors[i].BookID)
		if err != nil {
			return debtors, fmt.Errorf(libraryServicePath, err)
		}

		debtors[i].UserName = user.Name
		debtors[i].BookTitle = book.Title
	}

	return debtors, nil
}

func (s *libraryService) ListMetric(ctx context.Context, month int) (metric []*model.LibraryMetric, err error) {
	metric, err = s.borrS.ListMetric(ctx, month)
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
