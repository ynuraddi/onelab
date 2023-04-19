package model

type Transaction struct {
	UUID   string `json:"uuid"`
	UserID int    `json:"user_id"`
	BookID int    `json:"book_id"`
	Price  int    `json:"price"`
}
