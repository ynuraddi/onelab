package model

import "time"

type Debtor struct {
	BorrowID   int       `json:"borrow_id"   gorm:"column:id"`
	BorrowDate time.Time `json:"borrow_date" gorm:"column:borrow_date"`
	UserID     int       `json:"user_id"     gorm:"column:user_id"`
	UserName   string    `json:"user_name"   gorm:"column:name"`
	BookID     int       `json:"book_id"     gorm:"column:book_id"`
	BookTitle  string    `json:"book_name"   gorm:"column:title"`
}

func (Debtor) TableName() string {
	return "debtors"
}
