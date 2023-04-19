package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"app/model"

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

const userRepositoryPath = `userRepository: %w`

func (r *userRepository) Create(ctx context.Context, u model.CreateUserRq) error {
	err := r.db.WithContext(ctx).Create(&model.User{
		Login:    u.Login,
		Name:     u.Name,
		Password: u.Password,
		IsActive: false,
		Version:  1,
	}).Error
	// fmt.Println(err)
	// fmt.Println(errors.Is(err, gorm.ErrDuplicatedKey))

	// WARNING почему то не хэндлиться ошибка
	// костыль
	if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "duplicate key value") {
		return fmt.Errorf(userRepositoryPath, model.ErrUserIsAlreadyExist)
	} else if err != nil {
		return fmt.Errorf(userRepositoryPath, err)
	}

	return nil
}

func (r *userRepository) Get(ctx context.Context, id int) (u model.User, err error) {
	err = r.db.WithContext(ctx).
		First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, fmt.Errorf(userRepositoryPath, model.ErrUserIsNotExist)
	} else if err != nil {
		return u, fmt.Errorf(userRepositoryPath, err)
	}

	return u, nil
}

func (r *userRepository) Update(ctx context.Context, u model.User) error {
	err := r.db.WithContext(ctx).
		Where("id = ? and version = ?", u.ID, u.Version).
		Updates(model.User{
			Name:    u.Name,
			Login:   u.Login,
			Version: u.Version + 1,
		}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf(userRepositoryPath, model.ErrEditConflict)
	} else if err != nil {
		return fmt.Errorf(userRepositoryPath, err)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).
		Delete(&model.User{}, id).
		Error; err != nil {
		return fmt.Errorf(userRepositoryPath, err)
	}

	return nil
}

func (r *userRepository) GetByLogin(ctx context.Context, login string) (u model.User, err error) {
	err = r.db.WithContext(ctx).
		Table(model.User{}.TableName()).
		Where("login = ?", login).
		First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, fmt.Errorf(userRepositoryPath, model.ErrUserIsNotExist)
	}

	return u, err
}

func (r *userRepository) IsVerified(ctx context.Context, login string) (isActive bool, err error) {
	// context создал отдельно поскольку боюсь что мы прикрепим таймаут к основному котексту echo, и когда он будет обрабатывать страницу основную то случиться таймаут
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.db.WithContext(context).
		Table(model.User{}.TableName()).
		Select("is_active").
		Where("login = ?", login).
		First(&isActive).Error

	select {
	case <-context.Done():
		return false, fmt.Errorf(userRepositoryPath, model.ErrContextExceed)
	default:
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return isActive, fmt.Errorf(userRepositoryPath, model.ErrUserIsNotExist)
	} else if err != nil {
		return isActive, fmt.Errorf(userRepositoryPath, err)
	}

	return isActive, nil
}
