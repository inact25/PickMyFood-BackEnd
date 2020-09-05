package productCategoryRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type ProductCategoryRepo interface {
	AddProductCategory(productCategory *models.ProductCategory) error
	GetProductCategoryByID(id string) (*models.ProductCategory, error)
	GetAllProductCategory() ([]*models.ProductCategory, error)
	UpdateProductCategory(id string, productCategory *models.ProductCategory) error
	DeleteProductCategory(id string) error
}
