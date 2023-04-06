package model

type Book struct {
	ID     int    `json:"id"     gorm:"column:book_id"  `
	Name   string `json:"name"   gorm:"column:name"     `
	Author string `json:"author" gorm:"column:author"    `
}

func (Book) TableName() string {
	return "books"
}
