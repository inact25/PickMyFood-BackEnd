package orderUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	orderRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/order"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type OrderUsecaseImpl struct {
	orderRepo orderRepositories.OrderRepo
}

func InitOrderUseCaseImpl(order orderRepositories.OrderRepo) OrderUsecase {
	return &OrderUsecaseImpl{order}
}

// AddOrder usecase
func (o *OrderUsecaseImpl) AddOrder(order *models.Order) (*string, error) {
	println("Masuk Usecase")
	order.OrderCreated = utils.GetTimeNow()
	// for _, value := range order.SoldItems {

	// 	println("MASUK SINI")
	// 	product, _ := o.orderRepo.GetStock(value.ProductID)
	// 	if product.ProductStock < value.Qty {
	// 		message := fmt.Sprintf("%v Sementara Kosong", &product.ProductName)
	// 		return &message, nil
	// 	}
	// }
	// println("Masuk Order")
	_, err := o.orderRepo.AddOrder(order)
	if err != nil {
		return nil, err
	}
	return nil, nil
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
