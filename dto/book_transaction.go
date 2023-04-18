package dto

type BookTransaction struct {
	Book        string  `json:"book"`
	TotalAmount float64 `json:"total_amount"`
}
