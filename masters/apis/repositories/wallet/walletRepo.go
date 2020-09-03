package walletrepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type WalletRepo interface {
	GetWalletByID(id string) (*models.Wallet, error)
	TopUpWallet(topUP *models.TopUp, userID string) error
	// UpdatePoin(id string, user *models.User) error
}
