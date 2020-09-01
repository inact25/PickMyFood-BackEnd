package repositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type ProductsRepository interface {
	GetProducts() ([]*models.ProductModels, error)
	GetProductByID(ID string) ([]*models.ProductModels, error)

	GetProductsPrice() ([]*models.ProductPrice, error)

	GetProductsCategory() ([]*models.ProductCategory, error)
}
