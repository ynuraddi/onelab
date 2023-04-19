package model

import "time"

type BookBorrow struct {
	ID         int       `json:"borrow_id"   gorm:"column:id"          `
	UUID       string    `json:"uuid"        gorm:"column:uuid"        `
	BookID     int       `json:"book_id"     gorm:"column:book_id"     `
	UserID     int       `json:"user_id"     gorm:"column:user_id"     `
	BorrowDate time.Time `json:"borrow_date" gorm:"column:borrow_date" `
	ReturnDate time.Time `json:"return_date" gorm:"column:return_date" `
	Version    int       `json:"version"     gorm:"column:version"     `
}

func (BookBorrow) TableName() string {
	return "book_borrows"
}

type CreateBookBorrowRq struct {
	UUID       string    `json:"-"           gorm:"column:uuid"                                 `
	BookID     int       `json:"book_id"     gorm:"column:book_id"     validate:"required,min=1"`
	UserID     int       `json:"user_id"     gorm:"column:user_id"     validate:"required,min=1"`
	BorrowDate time.Time `json:"borrow_date" gorm:"column:borrow_date" validate:"required"      `
}

type UpdateBookBorrowRq struct {
	ID         int       `json:"-"           param:"id" validate:"required,min=1"`
	BookID     int       `json:"book_id"     `
	UserID     int       `json:"user_id"     `
	BorrowDate time.Time `json:"borrow_date" `
	ReturnDate time.Time `json:"return_date" `
}
