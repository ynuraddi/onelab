package model

import (
	"time"
)

type LibraryBorrowRq struct {
	UserLogin  string `json:"-"`
	BookTitle  string `json:"title" validate:"required,min=5"`
	BorrowDate string `json:"borrow_date"`
}

type LibraryBorrowRp struct {
	TransactionUUID string `json:"uuid"`
	Score           int    `json:"amount"`
}

type LibraryReturnRq struct {
	UserLogin string `json:"-" validate:"required"`
	BookTitle string `json:"title" validate:"required"`
}

type LibraryDebtor struct {
	BorrowID   int       `json:"borrow_id"   gorm:"column:id"`
	BorrowUUID string    `json:"uuid"        gorm:"column:uuid"`
	BorrowDate time.Time `json:"borrow_date" gorm:"column:borrow_date"`
	UserID     int       `json:"user_id"     gorm:"column:user_id"`
	UserName   string    `json:"user_name"   gorm:"column:name"`
	BookID     int       `json:"book_id"     gorm:"column:book_id"`
	BookTitle  string    `json:"book_name"   gorm:"column:title"`
}

type LibraryMetricRepo struct {
	UserID   int    `json:"user_id"        gorm:"column:user_id"`
	UserName string `json:"user_name"      gorm:"column:name"`
	Books    string `json:"books"          gorm:"column:books"`
}

type LibraryMetricUserBook struct {
	UserID   int      `json:"user_id"        gorm:"column:user_id"`
	UserName string   `json:"user_name"      gorm:"column:name"`
	Books    []string `json:"books"          gorm:"column:books"`
}

type LibraryMetricBookAmount struct {
	BookTitle string `json:"title"`
	Amount    int    `json:"amount"`
}
