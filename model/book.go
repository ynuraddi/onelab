package model

type Book struct {
	ID      int    `json:"id"      gorm:"column:id"     param:"id" validate:"required,min=1"`
	Title   string `json:"title"   gorm:"column:title"  `
	Author  string `json:"author"  gorm:"column:author" `
	Version int    `json:"version" gorm:"column:version"`
}

func (Book) TableName() string {
	return "books"
}

type CreateBookRq struct {
	Title  string `json:"title"  gorm:"column:title"  validate:"required,min=5"`
	Author string `json:"author" gorm:"column:author" validate:"required,min=5"`
}

type UpdateBookRq struct {
	ID     int    `json:"-"      param:"id"    validate:"required,min=1"`
	Title  string `json:"title"  `
	Author string `json:"author" `
}
