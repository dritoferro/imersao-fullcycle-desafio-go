package usecase

type TransactionDTOInput struct {
	AccountID string
	Amount    float64
}

type TransactionDTOOutput struct {
	ID        string
	AccountID string
	Amount    float64
}
