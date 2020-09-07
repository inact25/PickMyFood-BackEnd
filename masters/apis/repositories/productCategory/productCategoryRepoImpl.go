package productCategoryRepositories

import (
	"database/sql"
	"errors"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type ProductCategoryRepoImpl struct {
	db *sql.DB
}

func InitProductCategoryRepoImpl(db *sql.DB) ProductCategoryRepo {
	return &ProductCategoryRepoImpl{db: db}
}

func (pc *ProductCategoryRepoImpl) AddProductCategory(productCategory *models.ProductCategory) error {
	productCategoryID := guuid.New()
	tx, err := pc.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_PRODUCT_CATEGORY)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(productCategoryID, productCategory.ProductCategoryName); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

//GetProductByID
func (pc *ProductCategoryRepoImpl) GetProductCategoryByID(id string) (*models.ProductCategory, error) {
	stmt, err := pc.db.Prepare(utils.SELECT_PRODUCT_CATEGORY_BY_ID)
	productCategory := models.ProductCategory{}
	if err != nil {
		return &productCategory, err
	}
	errQuery := stmt.QueryRow(id).Scan(&productCategory.ProductCategoryID, &productCategory.ProductCategoryName)

	if errQuery != nil {
		return &productCategory, err
	}

	defer stmt.Close()
	return &productCategory, nil
}

// //GetAllProduct for admin
func (pc *ProductCategoryRepoImpl) GetAllProductCategory() ([]*models.ProductCategory, error) {
	stmt, err := pc.db.Prepare(utils.SELECT_ALL_PRODUCT_CATEGORY)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	listProductCategory := []*models.ProductCategory{}
	for rows.Next() {
		productCategory := models.ProductCategory{}
		err := rows.Scan(&productCategory.ProductCategoryID, &productCategory.ProductCategoryName)
		if err != nil {
			return nil, err
		}
		listProductCategory = append(listProductCategory, &productCategory)
	}
	return listProductCategory, nil
}

// // Update Store for profil
func (pc *ProductCategoryRepoImpl) UpdateProductCategory(id string, productCategory *models.ProductCategory) error {
	tx, err := pc.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_PRODUCT_CATEGORY)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(productCategory.ProductCategoryName, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// Delete Product Category for admin
func (pc *ProductCategoryRepoImpl) DeleteProductCategory(productCategoryID string) error {
	tx, err := pc.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_PRODUCT_CATEGORY)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(productCategoryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, store id tidak di temukan")
	}

	return tx.Commit()
}
