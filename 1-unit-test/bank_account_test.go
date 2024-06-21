package bank

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeposit(t *testing.T) {
	repository := &AccountRepositoryMock{}
	repository.On("SaveAccount", 105.0).Return(nil)
	repository.On("SaveAccount", 5.0).Return(errors.New("deposit should be greater than current balance"))
	repository.On("SaveAccount", mock.Anything).Return(nil)

	balance, _ := Deposit(100.0, 5.0, repository)
	assert.Equal(t, 105.0, balance)

	_, err := Deposit(0, 5.0, repository)
	assert.Error(t, err, "deposit should be greater than current balance")

	// balance, _ = Deposit(100.0, 100.0, repository)
	// assert.Equal(t, 200.0, balance)

}

func TestWithdraw(t *testing.T) {
	balance, _ := Withdraw(5.0, 100.0)
	assert.Equal(t, 95.0, balance)

	_, err := Withdraw(50.0, 49.99)
	assert.Error(t, err, "balance is not enough")
}
