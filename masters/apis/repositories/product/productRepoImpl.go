package productRepositories

import (
	"database/sql"
	"errors"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func InitProductRepoImpl(db *sql.DB) ProductRepo {
	return &ProductRepoImpl{db: db}
}

func (p *ProductRepoImpl) AddProduct(storeID string, product *models.Product) error {
	productID := guuid.New()
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_PRODUCT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(productID, storeID, product.ProductName, product.ProductCategory.ProductCategoryID, product.ProductStock, product.ProductStatus); err != nil {
		tx.Rollback()
		return err
	}

	productPriceID := guuid.New()
	stmt, err = tx.Prepare(utils.INSERT_PRODUCT_PRICE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(productPriceID, productID, product.ProductPrice.Price, product.ProductPrice.DateModified); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (p *ProductRepoImpl) GetProductByID(id string) (*models.Product, error) {
	stmt, err := p.db.Prepare(utils.SELECT_PRODUCT_BY_ID)
	product := models.Product{}
	if err != nil {
		return &product, err
	}
	errQuery := stmt.QueryRow(id).Scan(&product.ProductID, &product.ProductName, &product.ProductStock, &product.ProductStatus, &product.ProductCategory.ProductCategoryID, &product.ProductCategory.ProductCategoryName, &product.ProductPrice.Price, &product.ProductPrice.DateModified)

	if errQuery != nil {
		return &product, err
	}

	defer stmt.Close()
	return &product, nil
}

//get for customer
func (p *ProductRepoImpl) GetAllProductByStore(storeID string) ([]*models.Product, error) {
	stmt, err := p.db.Prepare(utils.SELECT_ALL_PRODUCT_BY_STORE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	listProduct := []*models.Product{}
	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.ProductCategory.ProductCategoryName, &product.ProductPrice.Price, &product.ProductStatus, &product.ProductPrice.DateModified)
		if err != nil {
			return nil, err
		}
		listProduct = append(listProduct, &product)
	}
	return listProduct, nil
}

func (p *ProductRepoImpl) UpdateProductWithPrice(id string, product *models.Product) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_PRODUCT_WITH_PRICE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(p, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	productPriceID := guuid.New()
	stmt, err = tx.Prepare(utils.INSERT_PRODUCT_PRICE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(productPriceID, id, product.ProductPrice.Price, product.ProductPrice.DateModified); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (p *ProductRepoImpl) DeleteProduct(id string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_PRODUCT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, product id tidak di temukan")
	}

	return tx.Commit()
}
