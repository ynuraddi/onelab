package model

type User struct {
	ID       int    `json:"id" param:"id" gorm:"column:user_id"`
	Name     string `json:"name" gorm:"column:name"`
	Login    string `json:"login" gorm:"column:login"`
	Password string `json:"-" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
