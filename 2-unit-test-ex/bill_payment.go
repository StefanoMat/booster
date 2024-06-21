package payment

import "errors"

type Repository interface {
	SavePayment(amount float64) error
	GetBill(id int) (float64, error)
}

type PaymentMethod interface {
	Pay(amount float64) error
}

func PayBill(id int, paymentMethod PaymentMethod, repository Repository) (float64, error) {
	amount, err := repository.GetBill(id)
	if err != nil {
		return 0.0, err
	}
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}

	err = paymentMethod.Pay(amount)
	if err != nil {
		return 0.0, err
	}
	err = repository.SavePayment(amount)
	if err != nil {
		return 0.0, err
	}
	return amount, nil
}
