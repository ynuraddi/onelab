package model

type User struct {
	ID       int    `json:"user_id"    gorm:"column:user_id"  `
	Name     string `json:"user_name"  gorm:"column:name"     `
	Login    string `json:"user_login" gorm:"column:login"    `
	Password string `json:"-"          gorm:"column:password" `
}

func (User) TableName() string {
	return "users"
}
