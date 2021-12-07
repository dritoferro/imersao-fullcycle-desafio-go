package usecase

import (
	"fmt"
	"github.com/dritoferro/imersao-fullcycle-desafio-go/entity"
	"time"
)

type ProcessTransaction struct {
	Repository entity.TransactionRepository
	counterIds int
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository, counterIds: 1}
}

func (p ProcessTransaction) PersistTransaction(input TransactionDTOInput) (TransactionDTOOutput, error) {
	transaction := entity.NewTransaction()

	transaction.ID = fmt.Sprint(p.counterIds)
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	err := p.Repository.Save(transaction.ID, transaction.AccountID, transaction.Amount, time.Date(2021, 12, 07, 20, 00, 00, 000, time.Local))
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	p.counterIds++

	return TransactionDTOOutput{
		ID:        transaction.ID,
		AccountID: transaction.AccountID,
		Amount:    transaction.Amount,
	}, nil
}
