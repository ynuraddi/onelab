package model

import (
	"time"
)

type LibraryBorrowRq struct {
	UserLogin string `json:"-"`
	BookTitle string `json:"title" validate:"required,min=5"`
	RentTerm  int    `json:"rent_term" validate:"required,min=1"`
}

type LibraryBorrowRp struct {
	TransactionUUID string `json:"transaction_uuid"`
	Score           int    `json:"score"`
}

type LibraryReturnRq struct {
	UserLogin string `json:"login" validate:"required"`
	BookTitle string `json:"title" validate:"required"`
}

type LibraryDebtor struct {
	BorrowID   int       `json:"borrow_id"   gorm:"column:id"`
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

type LibraryMetric struct {
	UserID   int      `json:"user_id"        gorm:"column:user_id"`
	UserName string   `json:"user_name"      gorm:"column:name"`
	Books    []string `json:"books"          gorm:"column:books"`
}
