package postgres

import (
	"context"
	"errors"
	"fmt"

	"app/model"

	"gorm.io/gorm"
)

type bookBorrowRepository struct {
	db *gorm.DB
}

func NewBookBorrowRepository(db *gorm.DB) *bookBorrowRepository {
	return &bookBorrowRepository{
		db: db,
	}
}

const bookBorrowRepositoryPath = `bookBorrowRepository: %w`

func (r *bookBorrowRepository) Create(ctx context.Context, record model.CreateBookBorrowRq) error {
	if err := r.db.WithContext(ctx).
		Select("uuid", "book_id", "user_id", "borrow_date", "version").
		Create(&model.BookBorrow{
			UUID:       record.UUID,
			BookID:     record.BookID,
			UserID:     record.UserID,
			BorrowDate: record.BorrowDate,
			Version:    1,
		}).
		Error; err != nil {
		return fmt.Errorf(bookBorrowRepositoryPath, err)
	}

	return nil
}

func (r *bookBorrowRepository) Get(ctx context.Context, id int) (b model.BookBorrow, err error) {
	err = r.db.WithContext(ctx).
		First(&b, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return b, fmt.Errorf(bookBorrowRepositoryPath, model.ErrBookBorrowIsNotExist)
	} else if err != nil {
		return b, fmt.Errorf(bookBorrowRepositoryPath, err)
	}

	return b, nil
}

func (r *bookBorrowRepository) Update(ctx context.Context, record model.BookBorrow) error {
	err := r.db.WithContext(ctx).
		Where("id = ? and version = ?", record.ID, record.Version).
		Updates(model.BookBorrow{
			BookID:     record.BookID,
			UserID:     record.UserID,
			BorrowDate: record.BorrowDate,
			ReturnDate: record.ReturnDate,
			Version:    record.Version + 1,
		}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf(bookBorrowRepositoryPath, model.ErrEditConflict)
	} else if err != nil {
		return fmt.Errorf(bookBorrowRepositoryPath, err)
	}

	return nil
}

func (r *bookBorrowRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Delete(&model.BookBorrow{}, id).
		Error; err != nil {
		return fmt.Errorf(bookBorrowRepositoryPath, err)
	}

	return nil
}

func (r *bookBorrowRepository) GetDebtors(ctx context.Context) (debtors []*model.Debtor, err error) {
	err = r.db.WithContext(ctx).Model(&model.BookBorrow{}).Where("return_date is null").Find(&debtors).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
		return debtors, fmt.Errorf(bookBorrowRepositoryPath, err)
	} else if err != nil {
		return debtors, fmt.Errorf(bookBorrowRepositoryPath, err)
	}

	return debtors, nil
}

func (r *bookBorrowRepository) GetMetric(ctx context.Context, month int) (metric []*model.Metric, err error) {
	err = r.db.WithContext(ctx).Model(&model.BookBorrow{}).
		Select("user_id, array_agg(title) as books").
		Where("extract(month from borrow_date) = ?", month).
		Group("user_id").
		Find(&metric).Error

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
		return metric, fmt.Errorf(bookBorrowRepositoryPath, err)
	} else if err != nil {
		return metric, fmt.Errorf(bookBorrowRepositoryPath, err)
	}

	return metric, nil
}

// func (r *bookBorrowRepository) ListDebtors(ctx context.Context) (debtors []*model.BookBorrowDebtorRp, err error) {
// 	bookBorrowTable := model.BookBorrow{}.TableName()
// 	// userTable := model.User{}.TableName()
// 	// bookTable := model.Book{}.TableName()

// 	err = r.db.WithContext(ctx).
// 		Table(bookBorrowTable).
// 		Joins("left join users on book_borrows.user_id = users.id").
// 		Joins("left join books on book_borrows.book_id = books.id").
// 		Select("*").
// 		Where("return_date is null").
// 		Find(&debtors).Error

// 	if errors.Is(err, gorm.ErrEmptySlice) || errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, fmt.Errorf(bookBorrowRepositoryPath, model.ErrBookBorrowIsNotExist)
// 	} else if err != nil {
// 		return nil, fmt.Errorf(bookBorrowRepositoryPath, err)
// 	}

// 	return debtors, nil
// }

// func (r *bookBorrowRepository) ListMetric(ctx context.Context, month int) (metric []*model.BookBorrowMetricRp, err error) {
// 	bookBorrowTable := model.BookBorrow{}.TableName()
// 	// userTable := model.User{}.TableName()

// 	err = r.db.WithContext(ctx).
// 		Table(bookBorrowTable).
// 		Joins("left join users on book_borrows.user_id = users.id").
// 		Select("user_id, name, count(borrow_date) as amount").
// 		Where("extract(month from borrow_date) = ?", month).
// 		Group("user_id, name").
// 		Find(&metric).Error

// 	if errors.Is(err, gorm.ErrEmptySlice) || errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, fmt.Errorf(bookBorrowRepositoryPath, model.ErrBookBorrowIsNotExist)
// 	} else if err != nil {
// 		return nil, fmt.Errorf(bookBorrowRepositoryPath, err)
// 	}

// 	return metric, nil
// }
