package entity

import "time"

type TransactionRepository interface {
	Save(id string, accountId string, amount float64, createdAt time.Time) error
}
