package usecase

import (
	mock_entity "github.com/dritoferro/imersao-fullcycle-desafio-go/entity/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSaveTransactionSuccessfully(t *testing.T) {
	input := TransactionDTOInput{
		AccountID: "123",
		Amount:    250,
	}

	outputExpected := TransactionDTOOutput{
		ID:        "1",
		AccountID: "123",
		Amount:    250,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)

	repositoryMock.EXPECT().Save("1", input.AccountID, input.Amount, time.Date(2021, 12, 07, 20, 00, 00, 000, time.Local)).Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	ouput, err := usecase.PersistTransaction(input)

	assert.Nil(t, err)
	assert.Equal(t, outputExpected, ouput)
}
