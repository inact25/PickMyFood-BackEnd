package walletusecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type WalletUseCases interface {
	GetWalletByID(id string) (*models.Wallet, error)
	TopUpWallet(topUP *models.TopUp, userID string) error
	// UpdateWallet(id string, wallet *models.Wallet) error
	// UpdatePoin(id string, user *models.User) error
}
