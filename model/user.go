package model

type User struct {
	ID       int    `json:"id"        gorm:"column:id"       param:"id" validate:"required,min=1"`
	Name     string `json:"user_name" gorm:"column:name"     `
	Login    string `json:"login"     gorm:"column:login"    `
	Password string `json:"-"         gorm:"column:password" `
	IsActive bool   `json:"is_active" gorm:"column:is_active"`
	Version  int    `json:"version"   gorm:"column:version"`
}

func (User) TableName() string {
	return "users"
}

// type contextKey string

// var ContextUsername = contextKey("username")

type LogInRq struct {
	Login    string `json:"login"                              validate:"required,min=5"`
	Password string `json:"password"                           validate:"required,min=5"`
}

type CreateUserRq struct {
	Name     string `json:"user_name"  gorm:"column:name"      validate:"required,min=5"`
	Login    string `json:"login"      gorm:"column:login"     validate:"required,min=5"`
	Password string `json:"password"   gorm:"column:password"  validate:"required,min=5"`
}

type UpdateUserRq struct {
	ID    int    `json:"-"         param:"id"       validate:"required,min=1"`
	Name  string `json:"user_name" `
	Login string `json:"login"     `
}
