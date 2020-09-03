package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type ProductUsecaseImpl struct {
	productRepo repositories.ProductsRepo
}

func (s ProductUsecaseImpl) GetProducts() ([]*models.ProductModels, error) {
	products, err := s.productRepo.GetProducts()

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s ProductUsecaseImpl) GetProductByID(ID string) (*models.ProductModels, error) {
	products, err := s.productRepo.GetProductByID(ID)

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s ProductUsecaseImpl) PostProduct(d models.ProductModels) (*models.ProductModels, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.productRepo.PostProduct(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s ProductUsecaseImpl) UpdateProduct(ID string, data models.ProductModels) (*models.ProductModels, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.productRepo.UpdateProduct(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s ProductUsecaseImpl) DeleteProduct(ID string) (*models.ProductModels, error) {
	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	_, err := s.productRepo.GetProductByID(ID)
	if err != nil {
		return nil, err
	}

	result, err := s.productRepo.DeleteProduct(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s ProductUsecaseImpl) GetProductsPrice() ([]*models.ProductPrice, error) {
	productsPrice, err := s.productRepo.GetProductsPrice()
	if err != nil {
		return nil, err
	}
	return productsPrice, nil
}

func (s ProductUsecaseImpl) GetProductsCategory() ([]*models.ProductCategory, error) {
	productsCategory, err := s.productRepo.GetProductsCategory()
	if err != nil {
		return nil, err
	}
	return productsCategory, nil
}

func InitProductUsecase(productRepo repositories.ProductsRepo) ProductsUseCases {
	return &ProductUsecaseImpl{productRepo}
}
