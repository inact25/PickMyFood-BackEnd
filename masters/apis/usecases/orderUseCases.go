package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type OrderUseCases interface {
	GetOrders() ([]*models.OrderModels, error)
	GetOrderByID(ID string) (*models.OrderModels, error)
	PostOrder(d models.OrderModels) (*models.OrderModels, error)
	UpdateOrder(ID string, data models.OrderModels) (*models.OrderModels, error)
	DeleteOrder(ID string) (*models.OrderModels, error)
}
