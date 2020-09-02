package repositories

import (
	"database/sql"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func (s *ProductRepoImpl) GetProducts() ([]*models.ProductModels, error) {
	var products []*models.ProductModels
	query := "SELECT * FROM tb_product"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.ProductModels{}
		err := rows.Scan(&product.ProductID, &product.StoreID, &product.ProductName, &product.ProductCategoryID, &product.ProductStatus)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)

	}

	return products, nil
}

func (s *ProductRepoImpl) GetProductByID(ID string) ([]*models.ProductModels, error) {
	var products []*models.ProductModels
	query := "SELECT * FROM tb_store WHERE product_id = ?"
	rows, err := s.db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.ProductModels{}
		err := rows.Scan(&product.ProductID, &product.StoreID, &product.ProductName, &product.ProductCategoryID, &product.ProductStatus)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)

	}

	return products, nil
}

func (s *ProductRepoImpl) GetProductsPrice() ([]*models.ProductPrice, error) {
	var productsPrice []*models.ProductPrice
	query := "SELECT * FROM tb_product_price"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		productPrice := models.ProductPrice{}
		err := rows.Scan(&productPrice.ProductPriceID, &productPrice.ProductID, &productPrice.Price, &productPrice.DateModified)

		if err != nil {
			return nil, err
		}

		productsPrice = append(productsPrice, &productPrice)

	}

	return productsPrice, nil
}

func (s *ProductRepoImpl) GetProductsCategory() ([]*models.ProductCategory, error) {
	var productsCategory []*models.ProductCategory
	query := "SELECT * FROM tb_product_category"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		productCategory := models.ProductCategory{}
		err := rows.Scan(&productCategory.ProductCategoryID, &productCategory.ProductCategoryName)

		if err != nil {
			return nil, err
		}

		productsCategory = append(productsCategory, &productCategory)

	}

	return productsCategory, nil
}

func InitProductRepoImpl(db *sql.DB) ProductsRepo {
	return &ProductRepoImpl{db}

}
