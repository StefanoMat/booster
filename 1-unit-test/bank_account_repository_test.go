package bank

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type BankAccountRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

type BankAccountDepositEntity struct {
	Id            int
	BankAccountId int
	Amount        float64
}

func (suite *BankAccountRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("create table bank_account_deposit (id integer primary key autoincrement, bank_account_id integer NOT NULL, amount float NOT NULL)")
	suite.Db = db
}

func (suite *BankAccountRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(BankAccountRepositoryTestSuite))
}

func (suite *BankAccountRepositoryTestSuite) TestSaveDeposit_ShouldInsertTheAmount() {
	repo := NewBankAccountRepository(suite.Db)
	var account int = 1
	var amount float64 = 100.0
	err := repo.SaveDeposit(account, amount)
	suite.NoError(err)
	var depositEntity BankAccountDepositEntity
	err = suite.Db.QueryRow("select id, bank_account_id, amount from bank_account_deposit where bank_account_id = ?", account).
		Scan(&depositEntity.Id, &depositEntity.BankAccountId, &depositEntity.Amount)

	suite.NoError(err)
	suite.NotNil(depositEntity.Id)
	suite.Equal(account, depositEntity.BankAccountId)
	suite.Equal(amount, depositEntity.Amount)
}
