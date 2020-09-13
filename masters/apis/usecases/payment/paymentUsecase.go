package paymentUsecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type PaymentUsecase interface {
	PaymentWallet(payment *models.Payment) error
	UpdateTransaction(storeID, amount, orderID, userID string) error
	GetAllTransactionByStore(storeID string) ([]*models.Payment, error)
	GetAllTransactionByUser(userID string) ([]*models.Payment, error)
	GetTransactionByID(id string) (*models.Payment, error)
}
