package postgres

import (
	"context"
	"errors"
	"fmt"

	"app/model"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

const bookRepositoryPath = `bookRepository: %w`

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(ctx context.Context, b model.CreateBookRq) error {
	err := r.db.WithContext(ctx).
		Create(&model.Book{
			Title:   b.Title,
			Author:  b.Author,
			Price:   b.Price,
			Version: 1,
		}).Error
	if err != nil {
		return fmt.Errorf(bookRepositoryPath, err)
	}

	return nil
}

func (r *bookRepository) Get(ctx context.Context, id int) (b model.Book, err error) {
	err = r.db.WithContext(ctx).
		First(&b, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return b, fmt.Errorf(bookRepositoryPath, model.ErrBookIsNotExist)
	} else if err != nil {
		return b, fmt.Errorf(bookRepositoryPath, err)
	}

	return b, nil
}

func (r *bookRepository) Update(ctx context.Context, b model.Book) error {
	err := r.db.WithContext(ctx).
		Where("id = ? and version = ?", b.ID, b.Version).
		Updates(model.Book{
			Title:   b.Title,
			Author:  b.Author,
			Version: b.Version + 1,
		}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf(bookRepositoryPath, model.ErrEditConflict)
	} else if err != nil {
		return fmt.Errorf(bookRepositoryPath, err)
	}

	return nil
}

func (r *bookRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Delete(&model.Book{}, id).
		Error; err != nil {
		return fmt.Errorf(bookRepositoryPath, err)
	}

	return nil
}
