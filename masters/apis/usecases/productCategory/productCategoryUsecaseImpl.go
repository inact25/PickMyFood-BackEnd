package productCategoryUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	productCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/productCategory"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type ProductCategoryUsecaseImpl struct {
	ProductCategoryRepo productCategoryRepositories.ProductCategoryRepo
}

func InitProductCategoryUseCase(productCategory productCategoryRepositories.ProductCategoryRepo) ProductCategoryUsecase {
	return &ProductCategoryUsecaseImpl{productCategory}
}

// AddProduct usecase
func (pc *ProductCategoryUsecaseImpl) AddProductCategory(productCategory *models.ProductCategory) error {
	err := validation.CheckEmpty(productCategory)
	if err != nil {
		return err
	}
	error := pc.ProductCategoryRepo.AddProductCategory(productCategory)
	if error != nil {
		return error
	}
	return nil
}

// GetProductById
func (pc *ProductCategoryUsecaseImpl) GetProductCategoryByID(productCategoryID string) (*models.ProductCategory, error) {
	productCategory, err := pc.ProductCategoryRepo.GetProductCategoryByID(productCategoryID)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}

func (pc *ProductCategoryUsecaseImpl) GetAllProductCategory() ([]*models.ProductCategory, error) {
	listProductCategory, err := pc.ProductCategoryRepo.GetAllProductCategory()
	if err != nil {
		return nil, err
	}
	return listProductCategory, nil
}

func (pc *ProductCategoryUsecaseImpl) UpdateProductCategory(id string, productCategory *models.ProductCategory) error {
	err := validation.CheckEmpty(productCategory)
	if err != nil {
		return err
	}
	error := pc.ProductCategoryRepo.UpdateProductCategory(id, productCategory)
	if error != nil {
		return error
	}
	return nil
}

func (pc *ProductCategoryUsecaseImpl) DeleteProductCategory(productCategoryID string) error {
	err := pc.ProductCategoryRepo.DeleteProductCategory(productCategoryID)
	if err != nil {
		return err
	}
	return nil
}
