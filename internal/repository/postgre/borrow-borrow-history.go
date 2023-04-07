package postgre

import (
	"context"
	"fmt"

	"app/internal/model"

	"gorm.io/gorm"
)

type bookBorrowHistoryRepository struct {
	db *gorm.DB
}

func NewBookBorrowHistoryRepository(db *gorm.DB) *bookBorrowHistoryRepository {
	return &bookBorrowHistoryRepository{
		db: db,
	}
}

func (r *bookBorrowHistoryRepository) Create(ctx context.Context, b model.BookBorrowHistory) error {
	if err := r.db.WithContext(ctx).Create(&b).Error; err != nil {
		return fmt.Errorf("bookBorrowHistoryRepo(Create): %w", err)
	}

	return nil
}

func (r *bookBorrowHistoryRepository) Get(ctx context.Context, id int) (b model.BookBorrowHistory, err error) {
	if err = r.db.WithContext(ctx).First(&b).Error; err != nil {
		return b, fmt.Errorf("bookBorrowHistoryRepo(Get): %w", err)
	}

	return b, nil
}

func (r *bookBorrowHistoryRepository) ListDebtors(ctx context.Context) (dbts []*model.Debtor, err error) {
	if err = r.db.WithContext(ctx).Table("book_borrowing_history").Select("*").Joins("left join users using(user_id) left join books using(book_id)").Where("return_date = null").Find(&dbts).Error; err != nil {
		return dbts, fmt.Errorf("bookBorrowHistoryRepo(ListDebtors): %w", err)
	}

	return dbts, nil
}

func (r *bookBorrowHistoryRepository) BookRentalForMonth(ctx context.Context, month, year int) (urb []*model.UserRentalBooks, err error) {
	if err = r.db.WithContext(ctx).Table("book_borrowing_history").Joins("left join users using(user_id)").Select("user_id, name, count(*)").Where("month(borrow_date) = ?", month).Group("user_id, name").Find(&urb).Error; err != nil {
		return urb, fmt.Errorf("bookBorrowHistoryRepo(BookRentalForMonth): %w", err)
	}

	return urb, nil
}
