package paymentRepositories

import (
	"database/sql"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type PaymentRepoImpl struct {
	db *sql.DB
}

func InitPaymentRepoImpl(db *sql.DB) PaymentRepo {
	return &PaymentRepoImpl{db: db}
}

func (p *PaymentRepoImpl) PaymentWallet(transaction *models.Payment) error {
	transactionID := guuid.New()
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_TRANSACTION)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(transactionID, transaction.OrderID, transaction.UserID, transaction.Amount, transaction.TransactionCreated); err != nil {
		tx.Rollback()
		return err
	}

	stmt, err = tx.Prepare(utils.UPDATE_WALLET_AMOUNT_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(transaction.Amount, transaction.UserID); err != nil {
		tx.Rollback()
		return err
	}
	stmt, err = tx.Prepare(utils.UPDATE_ORDER_DETAIL_STATUS)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(transaction.OrderID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// transaction pick up
func (p *PaymentRepoImpl) UpdateTransaction(storeID, amount, orderID, userID string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.UPDATE_TRANSACTION_PICK)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(orderID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
