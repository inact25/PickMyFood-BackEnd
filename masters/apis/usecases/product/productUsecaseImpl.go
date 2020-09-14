package productUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	productRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/product"
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type ProductUsecaseImpl struct {
	productRepo productRepositories.ProductRepo
}

func InitProductUseCaseImpl(product productRepositories.ProductRepo) ProductUsecase {
	return &ProductUsecaseImpl{product}
}

// AddProduct usecase
func (p *ProductUsecaseImpl) AddProduct(storeID string, product *models.Product) error {
	product.ProductPrice.DateModified = utils.GetTimeNow()
	err := validation.CheckEmpty(product)
	if err != nil {
		return err
	}
	error := p.productRepo.AddProduct(storeID, product)
	if error != nil {
		return error
	}
	return nil
}

// GetProductById
func (p *ProductUsecaseImpl) GetProductByID(id string) (*models.Product, error) {
	product, err := p.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductUsecaseImpl) GetAllProductByStore(storeID string) ([]*models.Product, error) {
	listProduct, err := p.productRepo.GetAllProductByStore(storeID)
	if err != nil {
		return nil, err
	}
	return listProduct, nil
}

func (p *ProductUsecaseImpl) UpdateProductWithPrice(id string, product *models.Product) error {
	product.ProductPrice.DateModified = utils.GetTimeNow()
	err := validation.CheckEmpty(product)
	if err != nil {
		return err
	}
	error := p.productRepo.UpdateProductWithPrice(id, product)
	if error != nil {
		return error
	}
	return nil
}

func (p *ProductUsecaseImpl) DeleteProduct(productID string) error {
	err := p.productRepo.DeleteProduct(productID)
	if err != nil {
		return err
	}
	return nil
}
