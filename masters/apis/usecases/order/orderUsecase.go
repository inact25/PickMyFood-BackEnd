package orderUsecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type OrderUsecase interface {
	AddOrder(storeID string, order *models.Order) error
	GetOrderByID(orderID string) (*models.Order, error)
	GetAllOrderByStore(storeID string) ([]*models.Order, error)
	//jika order sudah di bayar
	UpdateOrderPaid(orderID string, order *models.Order) error
	//jika terjadi cancel order
	UpdateOrderCancel(orderID string, payment *models.Payment) error
}
