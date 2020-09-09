package paymentRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type PaymentRepo interface {
	PaymentWallet(transaction *models.Payment) error
	UpdateTransaction(storeID, amount, orderID, userID string) error
}
