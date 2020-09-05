package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type ProductsUseCases interface {
	GetProducts() ([]*models.ProductModels, error)
	GetProductByID(ID string) (*models.ProductModels, error)
	PostProduct(d models.ProductModels) (*models.ProductModels, error)
	UpdateProduct(ID string, data models.ProductModels) (*models.ProductModels, error)
	DeleteProduct(ID string) (*models.ProductModels, error)

	GetProductsPrice() ([]*models.ProductPrice, error)
	GetProductPriceByID(ID string) (*models.ProductPrice, error)
	PostProductPrice(d models.ProductPrice) (*models.ProductPrice, error)
	UpdateProductPrice(ID string, data models.ProductPrice) (*models.ProductPrice, error)
	DeleteProductPrice(ID string) (*models.ProductPrice, error)

	GetProductsCategory() ([]*models.ProductCategory, error)
	GetProductCategoryByID(ID string) (*models.ProductCategory, error)
	PostProductCategory(d models.ProductCategory) (*models.ProductCategory, error)
	UpdateProductCategory(ID string, data models.ProductCategory) (*models.ProductCategory, error)
	DeleteProductCategory(ID string) (*models.ProductCategory, error)
}
