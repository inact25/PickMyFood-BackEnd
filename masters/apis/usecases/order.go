package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type OrderUsecaseImpl struct {
	orderRepo repositories.OrderRepo
}

func (s OrderUsecaseImpl) GetOrders() ([]*models.OrderModels, error) {
	orders, err := s.orderRepo.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s OrderUsecaseImpl) GetOrderByID(ID string) (*models.OrderModels, error) {
	orders, err := s.orderRepo.GetOrderByID(ID)

	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s OrderUsecaseImpl) PostOrder(d models.OrderModels) (*models.OrderModels, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.orderRepo.PostOrder(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s OrderUsecaseImpl) UpdateOrder(ID string, data models.OrderModels) (*models.OrderModels, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.orderRepo.UpdateOrder(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s OrderUsecaseImpl) DeleteOrder(ID string) (*models.OrderModels, error) {
	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	_, err := s.orderRepo.GetOrderByID(ID)
	if err != nil {
		return nil, err
	}

	result, err := s.orderRepo.DeleteOrder(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func InitOrderUsecase(orderRepo repositories.OrderRepo) OrderUseCases {
	return &OrderUsecaseImpl{orderRepo}
}
