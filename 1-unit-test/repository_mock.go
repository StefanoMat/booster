package bank

import (
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) SaveDeposit(id int, amount float64) error {
	args := m.Called(id, amount)
	return args.Error(0)
}
