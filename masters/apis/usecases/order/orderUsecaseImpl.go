package orderUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	orderRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/order"
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type OrderUsecaseImpl struct {
	orderRepo orderRepositories.OrderRepo
}

func InitOrderUseCaseImpl(order orderRepositories.OrderRepo) OrderUsecase {
	return &OrderUsecaseImpl{order}
}

// AddOrder usecase
func (o *OrderUsecaseImpl) AddOrder(order *models.Order) (*models.Order, error) {
	order.OrderCreated = utils.GetTimeNow()
	err := validation.CheckEmpty(order)
	if err != nil {
		return nil, err
	}
	error, _ := o.orderRepo.AddOrder(order)
	if error != nil {
		return error, nil
	}
	return error, err
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

func (o *OrderUsecaseImpl) GetAllOrderByUser(userID string) ([]*models.Order, error) {
	listOrder, err := o.orderRepo.GetAllOrderByUser(userID)
	if err != nil {
		return nil, err
	}
	return listOrder, nil
}
