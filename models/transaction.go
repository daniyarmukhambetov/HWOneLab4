package models

type Transaction struct {
	ID        int64
	UserID    int
	Amount    float64
	CreatedAt string
	BookName  string
}
