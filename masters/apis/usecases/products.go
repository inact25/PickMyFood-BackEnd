package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
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

func (s ProductUsecaseImpl) GetProductByID(ID string) ([]*models.ProductModels, error) {
	products, err := s.productRepo.GetProductByID(ID)

	if err != nil {
		return nil, err
	}
	return products, nil
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
