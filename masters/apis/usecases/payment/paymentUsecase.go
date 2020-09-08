package paymentUsecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type PaymentUsecase interface {
	PaymentWallet(payment *models.Payment) error
}
