package model

type User struct {
	ID       int    `json:"user_id"    gorm:"column:user_id"  `
	Name     string `json:"name"       gorm:"column:user_name"`
	Login    string `json:"login"      gorm:"column:login"    `
	Password string `json:"-"          gorm:"column:password" `
}

func (User) TableName() string {
	return "users"
}

type contextKey string

var ContextUsername = contextKey("username")

type LoginUserRq struct {
	Name     string `json:"name"     validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=5"`
}

type UserCreateRq struct {
	Name     string `json:"name"     validate:"required,min=5"`
	Login    string `json:"login"    validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=5"`
}

type UserUpdateRq struct {
	ID    int    `param:"id"      validate:"required,min=1"`
	Name  string `json:"name"     `
	Login string `json:"login"    `
}
