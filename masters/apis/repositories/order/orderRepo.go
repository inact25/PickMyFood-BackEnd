package orderRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type OrderRepo interface {
	AddOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(orderID string) (*models.Order, error)
	GetAllOrderByStore(storeID string) ([]*models.Order, error)
	GetAllOrderByUser(userID string) ([]*models.Order, error)
	GetStock(productID string) (*models.Product, error)
}
