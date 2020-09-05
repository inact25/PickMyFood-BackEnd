package orderRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type OrderRepo interface {
	AddOrder(storeID string, order *models.Order) error
	GetOrderByID(orderID string) (*models.Order, error)
	GetAllOrderByStore(storeID string) ([]*models.Order, error)
	//jika order sudah di bayar
	UpdateOrderPaid(orderID string, order *models.Order) error
	//jika terjadi cancel order
	UpdateOrder(orderID string, order *models.Order) error
}
