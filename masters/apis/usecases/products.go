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

func (s ProductUsecaseImpl) GetProductPriceByID(ID string) (*models.ProductPrice, error) {
	products, err := s.productRepo.GetProductPriceByID(ID)

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s ProductUsecaseImpl) PostProductPrice(d models.ProductPrice) (*models.ProductPrice, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.productRepo.PostProductPrice(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s ProductUsecaseImpl) UpdateProductPrice(ID string, data models.ProductPrice) (*models.ProductPrice, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.productRepo.UpdateProductPrice(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s ProductUsecaseImpl) DeleteProductPrice(ID string) (*models.ProductPrice, error) {
	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	_, err := s.productRepo.GetProductPriceByID(ID)
	if err != nil {
		return nil, err
	}

	result, err := s.productRepo.DeleteProductPrice(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s ProductUsecaseImpl) GetProductsCategory() ([]*models.ProductCategory, error) {
	productsCategory, err := s.productRepo.GetProductsCategory()
	if err != nil {
		return nil, err
	}
	return productsCategory, nil
}

func (s ProductUsecaseImpl) GetProductCategoryByID(ID string) (*models.ProductCategory, error) {
	products, err := s.productRepo.GetProductCategoryByID(ID)

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s ProductUsecaseImpl) PostProductCategory(d models.ProductCategory) (*models.ProductCategory, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.productRepo.PostProductCategory(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s ProductUsecaseImpl) UpdateProductCategory(ID string, data models.ProductCategory) (*models.ProductCategory, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.productRepo.UpdateProductCategory(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s ProductUsecaseImpl) DeleteProductCategory(ID string) (*models.ProductCategory, error) {
	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	_, err := s.productRepo.GetProductCategoryByID(ID)
	if err != nil {
		return nil, err
	}

	result, err := s.productRepo.DeleteProductCategory(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func InitProductUsecase(productRepo repositories.ProductsRepo) ProductsUseCases {
	return &ProductUsecaseImpl{productRepo}
}
