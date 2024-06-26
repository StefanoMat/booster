package bank

import (
	"database/sql"
	_ "database/sql"
)

type BankAccountRepository struct {
	Db *sql.DB
}

func NewBankAccountRepository(db *sql.DB) *BankAccountRepository {
	return &BankAccountRepository{Db: db}
}

func (r *BankAccountRepository) SaveDeposit(id int, amount float64) error {
	stmt, err := r.Db.Prepare("insert into bank_account_deposit (bank_account_id, amount) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, amount)
	if err != nil {
		return err
	}
	return nil
}
