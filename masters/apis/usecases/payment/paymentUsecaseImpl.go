package paymentUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	paymentRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/payment"
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type PaymentUsecaseImpl struct {
	PaymentRepo paymentRepositories.PaymentRepo
}

func InitPaymentUseCaseImpl(payment paymentRepositories.PaymentRepo) PaymentUsecase {
	return &PaymentUsecaseImpl{payment}
}

// AddPayment usecase
func (p *PaymentUsecaseImpl) PaymentWallet(payment *models.Payment) error {
	println("MASUK USECASE")
	payment.TransactionCreated = utils.GetTimeNow()
	err := validation.CheckEmpty(payment)
	if err != nil {
		return err
	}
	error := p.PaymentRepo.PaymentWallet(payment)
	if error != nil {
		return error
	}
	return nil
}
