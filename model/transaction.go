package model

type CreateTransactionRq struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
	Amount int `json:"amount"`
}

type CreateTransactionRp struct {
	UUID   string `json:"uuid"`
	Amount int    `json:"amount"`
}

type PayTransactionRq struct {
	UUID   string `json:"uuid"`
	Amount int    `json:"amount"`
}

type ErrTransactionRp struct {
	Err string `json:"error"`
}
