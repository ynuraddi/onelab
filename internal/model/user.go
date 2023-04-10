package model

type User struct {
	ID       int    `json:"user_id"    gorm:"column:user_id"  `
	Name     string `json:"name"  gorm:"column:user_name"     `
	Login    string `json:"login" gorm:"column:login"    `
	Password string `json:"-"          gorm:"column:password" `
}

func (User) TableName() string {
	return "users"
}
