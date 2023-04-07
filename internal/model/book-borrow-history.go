package model

import "time"

type BookBorrowHistory struct {
	ID         int       `json:"borrow_id"   gorm:"column:borrow_id"`
	BookID     int       `json:"book_id"     gorm:"column:book_id"`
	UserID     int       `json:"user_id"     gorm:"column:user_id"`
	BorrowDate time.Time `json:"borrow_date" gorm:"column:borrow_date"`
	ReturnDate time.Time `json:"return_date" gorm:"column:return_date"`
}

func (BookBorrowHistory) TableName() string {
	return "book_borrowing_history"
}

type Debtor struct {
	UserID     int       `json:"user_id"`
	UserName   string    `json:"user_name"`
	BookID     int       `json:"book_id"`
	BookTitle  string    `json:"book_name"`
	BorrowDate time.Time `json:"borrow_date"`
}

type UserRentalBooks struct {
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	CountBooks int    `json:"count_book"`
}
