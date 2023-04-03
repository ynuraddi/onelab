package repository

import (
	"fmt"
	"sync"

	"app/internal/model"
)

type userRepository struct {
	db    map[int]model.User
	count int
	mu    sync.RWMutex
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		db:    map[int]model.User{},
		count: 0,
		mu:    sync.RWMutex{},
	}
}

func (r *userRepository) Create(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.count == 3 {
		return fmt.Errorf("userRpository(Create): %w", model.ErrUserOverflow)
	}

	r.count++
	user.ID = r.count
	r.db[r.count] = user

	return nil
}

func (r *userRepository) Get(id int) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exist := r.db[id]
	if !exist {
		return model.User{}, fmt.Errorf("userRepository(Get): %w", model.ErrUserNotExists)
	}

	return user, nil
}
