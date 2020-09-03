package repositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type ProductsRepo interface {
	GetProducts() ([]*models.ProductModels, error)
	GetProductByID(ID string) (*models.ProductModels, error)
	PostProduct(d models.ProductModels) (*models.ProductModels, error)
	UpdateProduct(ID string, data models.ProductModels) (*models.ProductModels, error)
	DeleteProduct(ID string) (*models.ProductModels, error)

	GetProductsPrice() ([]*models.ProductPrice, error)

	GetProductsCategory() ([]*models.ProductCategory, error)
}
