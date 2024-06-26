package bank

import "errors"

type Repository interface {
	SaveDeposit(accountId int, amount float64) error
}

func Deposit(id int, amount float64, balance float64, repository Repository) (float64, error) {
	balance += amount
	err := repository.SaveDeposit(id, amount)
	if err != nil {
		return 0.0, err
	}
	return balance, nil
}

func Withdraw(amount float64, balance float64) (float64, error) {
	if amount > balance {
		return 0.0, errors.New("balance is not enough")
	}
	balance -= amount
	return balance, nil
}
