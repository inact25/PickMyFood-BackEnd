package storeCategoryRepositories

import (
	"database/sql"
	"errors"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type StoreCategoryRepoImpl struct {
	db *sql.DB
}

func InitStoreCategoryRepoImpl(db *sql.DB) StoreCategoryRepo {
	return &StoreCategoryRepoImpl{db: db}
}

func (sc *StoreCategoryRepoImpl) AddStoreCategory(store *models.StoreCategory) error {
	storeCategoryID := guuid.New()
	tx, err := sc.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_STORE_CATEGORY)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(storeCategoryID, store.StoreCategoryName); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

//GetStoreByID
func (sc *StoreCategoryRepoImpl) GetStoreCategoryByID(id string) (*models.StoreCategory, error) {
	stmt, err := sc.db.Prepare(utils.SELECT_STORE_CATEGORY_BY_ID)
	storeCategory := models.StoreCategory{}
	if err != nil {
		return &storeCategory, err
	}
	errQuery := stmt.QueryRow(id).Scan(&storeCategory.StoreCategoryID, &storeCategory.StoreCategoryName)

	if errQuery != nil {
		return &storeCategory, err
	}

	defer stmt.Close()
	return &storeCategory, nil
}

// //GetAllUser for admin
func (sc *StoreCategoryRepoImpl) GetAllStoreCategory() ([]*models.StoreCategory, error) {
	stmt, err := sc.db.Prepare(utils.SELECT_ALL_STORE_CATEGORY)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	listStoreCategory := []*models.StoreCategory{}
	for rows.Next() {
		storeCategory := models.StoreCategory{}
		err := rows.Scan(&storeCategory.StoreCategoryID, &storeCategory.StoreCategoryName)
		if err != nil {
			return nil, err
		}
		listStoreCategory = append(listStoreCategory, &storeCategory)
	}
	return listStoreCategory, nil
}

// // Update Store for profil
func (sc *StoreCategoryRepoImpl) UpdateStoreCategory(id string, storeCategory *models.StoreCategory) error {
	tx, err := sc.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_STORE_CATEGORY)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(storeCategory.StoreCategoryName, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// Delete Store Category for admin
func (sc *StoreCategoryRepoImpl) DeleteStoreCategory(storeCategoryID string) error {
	tx, err := sc.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_STORE_CATEGORY)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(storeCategoryID)
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
