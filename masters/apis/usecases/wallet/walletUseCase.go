package walletusecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type WalletUseCases interface {
	GetWalletByID(id string) (*models.Wallet, error)
	TopUpWallet(topUP *models.TopUp, userID string) error
	UpdateAmountWallet(wallet *models.Wallet, userID string) error
	GetAllTopUp() ([]*models.TopUp, error)
	GetAllTopUpByUser(userID string) ([]*models.TopUp, error)
}
