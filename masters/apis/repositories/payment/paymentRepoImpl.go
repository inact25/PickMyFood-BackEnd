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
	stmt, err = tx.Prepare(utils.UPDATE_POIN_USER_AFTER_PAYMENT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(userID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (p *PaymentRepoImpl) GetAllTransactionByStore(storeID string) ([]*models.Payment, error) {
	stmt, err := p.db.Prepare(utils.SELECT_ALL_TRANSACTION_BY_STORE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(storeID)
	if err != nil {
		return nil, err
	}

	listTransaction := []*models.Payment{}
	for rows.Next() {
		transaction := models.Payment{}
		err := rows.Scan(&transaction.TransactionID, &transaction.OrderID, &transaction.UserID, &transaction.UserFirstName, &transaction.Amount, &transaction.TransactionCreated, &transaction.TransactionStatus)
		if err != nil {
			return nil, err
		}
		listTransaction = append(listTransaction, &transaction)
	}
	return listTransaction, nil
}

func (p *PaymentRepoImpl) GetAllTransactionByUser(userID string) ([]*models.Payment, error) {
	stmt, err := p.db.Prepare(utils.SELECT_ALL_TRANSACTION_BY_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}

	listTransaction := []*models.Payment{}
	for rows.Next() {
		transaction := models.Payment{}
		err := rows.Scan(&transaction.TransactionID, &transaction.OrderID, &transaction.UserID, &transaction.UserFirstName, &transaction.Amount, &transaction.TransactionCreated, &transaction.TransactionStatus)
		if err != nil {
			return nil, err
		}
		listTransaction = append(listTransaction, &transaction)
	}
	return listTransaction, nil
}
func (p *PaymentRepoImpl) GetTransactionByID(id string) (*models.Payment, error) {
	stmt, err := p.db.Prepare(utils.SELECT_TRANSACTION_BY_ID)
	transaction := models.Payment{}
	if err != nil {
		return &transaction, err
	}
	errQuery := stmt.QueryRow(id).Scan(&transaction.TransactionID, &transaction.OrderID, &transaction.UserID, &transaction.UserFirstName, &transaction.Amount, &transaction.TransactionCreated, &transaction.TransactionStatus)

	if errQuery != nil {
		return &transaction, err
	}

	defer stmt.Close()
	return &transaction, nil
}

func (p *PaymentRepoImpl) GetValidation(orderID, storeID string) bool {
	stmt, err := p.db.Prepare(utils.SELECT_VALIDATION_ORDER)
	order := models.Order{}
	if err != nil {
		println(err)
		return false
	}
	errQuery := stmt.QueryRow(orderID, storeID).Scan(&order.OrderID, &order.OrderCreated, &order.StoreID)

	if errQuery != nil {
		println(errQuery)
		return false
	}
	defer stmt.Close()
	return true
}
