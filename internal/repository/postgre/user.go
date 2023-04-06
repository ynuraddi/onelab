package postgre

import (
	"context"
	"fmt"

	"app/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, u model.User) error {
	result := r.db.WithContext(ctx).Create(&u)
	if result.Error != nil {
		return fmt.Errorf("userRepository(Create): %w", result.Error)
	}

	return nil
}

func (r *userRepository) Get(ctx context.Context, id int) (u model.User, err error) {
	result := r.db.WithContext(ctx).First(&u, id)
	if result.Error != nil {
		return u, fmt.Errorf("userRepository(Get): %w", result.Error)
	}

	return u, nil
}

func (r *userRepository) Update(ctx context.Context, u model.User) error {
	result := r.db.WithContext(ctx).Where("user_id = ?", u.ID).Updates(model.User{Name: u.Name, Login: u.Login})
	if result.Error != nil {
		return fmt.Errorf("userRepository(Update): %w", result.Error)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(model.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("userRepository(Delete): %w", result.Error)
	}

	return nil
}
