package productRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type ProductRepo interface {
	AddProduct(storeID string, product *models.Product) error
	GetProductByID(id string) (*models.Product, error)
	GetAllProductByStore(storeID string) ([]*models.Product, error)
	UpdateProductWithPrice(id string, product *models.Product) error
	DeleteProduct(id string) error
	GetProductNonAktif(storeID string) ([]*models.Product, error)
	ChangeActive(id string) error
}
