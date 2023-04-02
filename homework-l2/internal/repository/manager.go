package repository

type IUserRepository interface {
	Create()
	Get()
	Update()
	Delete()
}

type Repository struct {
	User IUserRepository
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserRepository(),
	}
}
