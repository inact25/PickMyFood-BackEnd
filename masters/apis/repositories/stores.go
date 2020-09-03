package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
)

type StoreRepoImpl struct {
	db *sql.DB
}

func (s *StoreRepoImpl) GetStores() ([]*models.StoreModels, error) {
	var stores []*models.StoreModels
	query := "SELECT * FROM tb_store"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		store := models.StoreModels{}
		err := rows.Scan(&store.StoreID, &store.StoreName, &store.StoreCategoryID, &store.StoreAddress, &store.StoreOwner, &store.StoreStatus, &store.StorePassword, &store.StoreImages)

		if err != nil {
			return nil, err
		}

		stores = append(stores, &store)

	}

	return stores, nil
}

func (s *StoreRepoImpl) GetStoreByID(ID string) (*models.StoreModels, error) {
	results := s.db.QueryRow("SELECT * FROM tb_store WHERE store_id = ?", ID)

	var d models.StoreModels
	err := results.Scan(&d.StoreID, &d.StoreName, &d.StoreCategoryID, &d.StoreAddress, &d.StoreOwner, &d.StoreStatus, &d.StorePassword, &d.StoreImages)
	if err != nil {
		return nil, errors.New("Menu ID Not Found")
	}

	return &d, nil
}

func (s *StoreRepoImpl) PostStore(d models.StoreModels) (*models.StoreModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO tb_store(store_id, store_name, store_category_id, store_address, store_owner, store_status, store_password, store_images) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.StoreID, d.StoreName, d.StoreCategoryID, d.StoreAddress, d.StoreOwner, d.StoreStatus, d.StorePassword, d.StoreImages)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return s.GetStoreByID(strconv.Itoa(int(lastInsertID)))
}

func (s *StoreRepoImpl) UpdateStore(ID string, data models.StoreModels) (*models.StoreModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE tb_store SET store_name=?, store_category_id=?, store_address=?, store_owner=?, store_status=?, store_password=?, store_images=? WHERE store_id=?`,
		data.StoreName, data.StoreCategoryID, data.StoreAddress, data.StoreOwner, data.StoreStatus, data.StorePassword, data.StoreImages, ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.GetStoreByID(ID)
}

func (s *StoreRepoImpl) DeleteStore(ID string) (*models.StoreModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM tb_store WHERE store_id = ?", ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return s.GetStoreByID(ID)

}

func (s *StoreRepoImpl) GetStoresCategory() ([]*models.StoreCategory, error) {
	var storesCategory []*models.StoreCategory
	query := "SELECT * FROM tb_store_category"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		storeCategory := models.StoreCategory{}
		err := rows.Scan(&storeCategory.StoreCategoryID, &storeCategory.StoreCategoryName)

		if err != nil {
			return nil, err
		}

		storesCategory = append(storesCategory, &storeCategory)

	}

	return storesCategory, nil
}

func InitStoreRepoImpl(db *sql.DB) StoresRepo {
	return &StoreRepoImpl{db}

}
