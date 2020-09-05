package repositories

import (
	"database/sql"

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

func (s *StoreRepoImpl) GetStoreByID(ID string) ([]*models.StoreModels, error) {
	var stores []*models.StoreModels
	query := "SELECT * FROM tb_store WHERE store_id = ?"
	rows, err := s.db.Query(query, ID)
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

func (s *StoreRepoImpl) DeleteStore(ID string) ([]*models.StoreModels, error) {
	var stores []*models.StoreModels
	query := "DELETE FROM tb_store WHERE store_id = ?"
	_, err := s.db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	// for rows.Next() {
	// 	store := models.StoreModels{}
	// 	err := rows.Scan(&store.StoreID, &store.StoreName, &store.StoreCategoryID, &store.StoreAddress, &store.StoreOwner, &store.StoreStatus, &store.StorePassword, &store.StoreImages)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	stores = append(stores, &store)

	// }

	return stores, nil
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

func InitStoreRepoImpl(db *sql.DB) StoresRepository {
	return &StoreRepoImpl{db}

}
