package orderUsecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type OrderUsecase interface {
	AddOrder(order *models.Order) (*string, error)
	GetOrderByID(orderID string) (*models.Order, error)
	GetAllOrderByStore(storeID string) ([]*models.Order, error)
	GetAllOrderByUser(userID string) ([]*models.Order, error)
}
