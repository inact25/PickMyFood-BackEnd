package walletrepositories

import (
	"database/sql"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type WalletRepoImpl struct {
	db *sql.DB
}

func InitWalletRepoImpl(db *sql.DB) WalletRepo {
	return &WalletRepoImpl{db: db}
}

// wallet user by id
func (w *WalletRepoImpl) GetWalletByID(id string) (*models.Wallet, error) {
	stmt, err := w.db.Prepare(utils.SELECT_WALLET_USER_ID)
	wallet := models.Wallet{}
	if err != nil {
		return &wallet, err
	}
	errQuery := stmt.QueryRow(id).Scan(&wallet.WalletID, &wallet.Amount, &wallet.UserID)

	if errQuery != nil {
		return &wallet, err
	}

	defer stmt.Close()
	return &wallet, nil
}

//TopUP wallet default unpaid
func (w *WalletRepoImpl) TopUpWallet(topUP *models.TopUp, userID string) error {
	topUpID := guuid.New()
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_TOP_UP)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(topUpID, topUP.TopUpAmount, userID, topUP.TopUpDate); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// update wallet
func (w *WalletRepoImpl) UpdateAmountWallet(wallet *models.Wallet, userID string) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_AMOUNT_WALLET)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(wallet.Amount, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	stmt, err = tx.Prepare(utils.UPDATE_STATUS_TOP_UP)

	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec("Paid", userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (w *WalletRepoImpl) GetAllTopUp() ([]*models.TopUp, error) {
	stmt, err := w.db.Prepare(utils.SELECT_ALL_TOP_UP)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	listTopUp := []*models.TopUp{}
	for rows.Next() {
		topUP := models.TopUp{}
		err := rows.Scan(&topUP.TopUpID, &topUP.TopUpAmount, &topUP.UserID, &topUP.UserFirstName, &topUP.TopUpDate, &topUP.TopUpStatus)
		if err != nil {
			return nil, err
		}
		listTopUp = append(listTopUp, &topUP)
	}
	return listTopUp, nil
}
func (w *WalletRepoImpl) GetAllTopUpByUser(userID string) ([]*models.TopUp, error) {
	stmt, err := w.db.Prepare(utils.SELECT_ALL_TOP_UP_BY_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	listTopUp := []*models.TopUp{}
	for rows.Next() {
		topUP := models.TopUp{}
		err := rows.Scan(&topUP.TopUpID, &topUP.TopUpAmount, &topUP.UserID, &topUP.UserFirstName, &topUP.TopUpDate, &topUP.TopUpStatus)
		if err != nil {
			return nil, err
		}
		listTopUp = append(listTopUp, &topUP)
	}
	return listTopUp, nil
}
