package orderUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	orderRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/order"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type OrderUsecaseImpl struct {
	orderRepo orderRepositories.OrderRepo
}

func InitOrderUseCaseImpl(order orderRepositories.OrderRepo) OrderUsecase {
	return &OrderUsecaseImpl{order}
}

// AddOrder usecase
func (o *OrderUsecaseImpl) AddOrder(storeID string, order *models.Order) error {
	err := validation.CheckEmpty(order)
	if err != nil {
		return err
	}
	error := o.orderRepo.AddOrder(storeID, order)
	if error != nil {
		return error
	}
	return nil
}

// GetOrderById
func (o *OrderUsecaseImpl) GetOrderByID(id string) (*models.Order, error) {
	order, err := o.orderRepo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderUsecaseImpl) GetAllOrderByStore(storeID string) ([]*models.Order, error) {
	listOrder, err := o.orderRepo.GetAllOrderByStore(storeID)
	if err != nil {
		return nil, err
	}
	return listOrder, nil
}

func (o *OrderUsecaseImpl) UpdateOrderPaid(id string, order *models.Order) error {
	err := validation.CheckEmpty(order)
	if err != nil {
		return err
	}
	error := o.orderRepo.UpdateOrderPaid(id, order)
	if error != nil {
		return error
	}
	return nil
}

func (o *OrderUsecaseImpl) UpdateOrderCancel(orderID string, payment *models.Payment) error {
	err := validation.CheckEmpty(payment)
	if err != nil {
		return err
	}
	error := o.orderRepo.UpdateOrderCancel(orderID, payment)
	if error != nil {
		return error
	}
	return nil
}
