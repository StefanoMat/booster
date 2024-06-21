package bank

import "errors"

type Repository interface {
	SaveAccount(balance float64) error
}

func Deposit(amount float64, balance float64, repository Repository) (float64, error) {
	balance += amount
	err := repository.SaveAccount(balance)
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
