package postgre

import (
	"context"
	"fmt"

	"app/internal/model"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(ctx context.Context, b model.Book) error {
	if err := r.db.WithContext(ctx).Create(&b).Error; err != nil {
		return fmt.Errorf("bookRepository(Create): %w", err)
	}

	return nil
}

func (r *bookRepository) Get(ctx context.Context, id int) (b model.Book, err error) {
	if err = r.db.WithContext(ctx).First(&b, id).Error; err != nil {
		return b, fmt.Errorf("bookRepository(Get): %w", err)
	}

	return b, nil
}
