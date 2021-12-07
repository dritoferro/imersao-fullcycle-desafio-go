package entity

import "time"

type Transaction struct {
	ID        string
	AccountID string
	Amount    float64
	createdAt time.Time
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
