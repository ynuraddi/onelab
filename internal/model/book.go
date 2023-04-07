package model

type Book struct {
	ID     int    `json:"book_id"     gorm:"column:book_id"  `
	Title  string `json:"book_title"   gorm:"column:title"     `
	Author string `json:"book_author" gorm:"column:author"    `
}

func (Book) TableName() string {
	return "books"
}
