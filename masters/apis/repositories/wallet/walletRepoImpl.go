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
	errQuery := stmt.QueryRow(id).Scan(&wallet.WalletID, &wallet.Amount, &wallet.UserID, &wallet.User.UserFirstName, &wallet.User.UserLastName, &wallet.User.UserAddress, &wallet.User.UserPhone, &wallet.User.UserPoin)

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
// func (w *WalletRepoImpl) UpdateWallet(id string, wallet *models.Wallet) error {
// 	tx, err := w.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	stmt, err := tx.Prepare(utils.UPDATE_AMOUNT_WALLET)
// 	defer stmt.Close()
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	_, err = stmt.Exec(wallet.Amount, id)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit()
// }

// // update poin
// func (w *WalletRepoImpl) UpdatePoin(id string, user *models.User) error {
// 	tx, err := w.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	stmt, err := tx.Prepare(utils.UPDATE_POIN_USER)
// 	defer stmt.Close()
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	_, err = stmt.Exec(user.UserPoin, id)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit()
// }
