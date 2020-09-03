package walletusecases

import (
	"strconv"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	walletrepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/wallet"
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type WalletUseCaseImpl struct {
	walletRepo walletrepositories.WalletRepo
}

func InitWalletUseCase(wallet walletrepositories.WalletRepo) WalletUseCases {
	return &WalletUseCaseImpl{wallet}
}

func (w *WalletUseCaseImpl) GetWalletByID(id string) (*models.Wallet, error) {
	wallet, err := w.walletRepo.GetWalletByID(id)
	if err != nil {
		return nil, err
	}
	return wallet, nil

}
func (w *WalletUseCaseImpl) TopUpWallet(topUP *models.TopUp, userID string) error {
	topUP.TopUpDate = utils.GetTimeNow()
	err := validation.CheckEmpty(topUP)
	if err != nil {
		return err
	}
	error := w.walletRepo.TopUpWallet(topUP, userID)
	if error != nil {
		return error
	}
	return nil
}

func (w *WalletUseCaseImpl) UpdateAmountWallet(wallet *models.Wallet, userID string) error {
	amount, _ := strconv.Atoi(wallet.Amount)
	println(amount)
	var totalAmount int = amount
	totalAmount = totalAmount + amount
	println(totalAmount)
	newTotalAmount := strconv.Itoa(totalAmount)
	println(newTotalAmount)

	wallet.Amount = newTotalAmount
	err := validation.CheckEmpty(wallet.Amount)
	if err != nil {
		return err
	}
	error := w.walletRepo.UpdateAmountWallet(wallet, userID)
	if error != nil {
		return error
	}
	return nil
}

// func (w *WalletUseCaseImpl) UpdatePoin(id string, user *models.User) error {
// 	err := validation.CheckEmpty(user.UserPoin)
// 	if err != nil {
// 		return err
// 	}
// 	error := w.walletRepo.UpdatePoin(id, user)
// 	if error != nil {
// 		return error
// 	}
// 	return nil
// }
