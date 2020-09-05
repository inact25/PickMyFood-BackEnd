package productRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type ProductRepo interface {
	AddProduct(storeID string, product *models.Product) error
	GetProductByID(id string) (*models.Product, error)
	GetAllProductByStore(storeID string) ([]*models.Product, error)
	// GetAllProductPrice() ([]*models.Product, error)
	UpdateProductWithPrice(id string, product *models.Product) error
	// UpdateProductPrice(productID string, productPrice *models.ProductPrice) error
	DeleteProduct(id string) error
}
