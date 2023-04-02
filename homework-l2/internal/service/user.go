package service

import "app/internal/repository"

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (r *userService) Create() {
	// validation
	// notification
	// ...
	r.repo.Create()
}

func (r *userService) Get() {
}

func (r *userService) Update() {
}

func (r *userService) Delete() {
}
