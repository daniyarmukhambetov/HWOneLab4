package models

type Transaction struct {
	ID        int64   `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
	BookName  string  `json:"book_name"`
}

type TransactionIn struct {
	UserID   int     `json:"user_id"`
	Amount   float64 `json:"amount"`
	BookName string  `json:"book_name"`
}
