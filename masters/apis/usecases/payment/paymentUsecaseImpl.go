package paymentUsecases

import (
	"errors"

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

func (p *PaymentUsecaseImpl) UpdateTransaction(storeID, amount, orderID, userID string) error {
	if !p.PaymentRepo.GetValidation(orderID, storeID) {
		println("MASUK SINI", p.PaymentRepo.GetValidation(orderID, storeID))
		return errors.New("Not Valid")
	}
	err := validation.CheckEmpty(storeID, amount, orderID, amount)
	if err != nil {
		return err
	}
	error := p.PaymentRepo.UpdateTransaction(storeID, amount, orderID, userID)
	if error != nil {
		return error
	}
	return nil
}
func (p *PaymentUsecaseImpl) GetAllTransactionByStore(storeID string) ([]*models.Payment, error) {
	listPayment, err := p.PaymentRepo.GetAllTransactionByStore(storeID)
	if err != nil {
		return nil, err
	}
	return listPayment, nil
}
func (p *PaymentUsecaseImpl) GetAllTransactionByUser(userID string) ([]*models.Payment, error) {
	listPayment, err := p.PaymentRepo.GetAllTransactionByUser(userID)
	if err != nil {
		return nil, err
	}
	return listPayment, nil
}
func (p *PaymentUsecaseImpl) GetTransactionByID(id string) (*models.Payment, error) {
	Payment, err := p.PaymentRepo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}
	return Payment, nil
}
