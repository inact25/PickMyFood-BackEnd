package storerepositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type StoreRepoImpl struct {
	db *sql.DB
}

func InitStoreRepoImpl(db *sql.DB) StoreRepo {
	return &StoreRepoImpl{db: db}
}

func (s *StoreRepoImpl) AddStore(store *models.Store) error {
	storeID := guuid.New()
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_STORE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(storeID, store.StoreName, store.StoreCategory.StoreCategoryID, store.StoreAddress, store.StoreOwner, store.StoreUsername, store.StorePassword); err != nil {
		tx.Rollback()
		return err
	}

	// stmt, err = tx.Prepare(utils.INSERT_STORE_CATEGORY)
	// defer stmt.Close()
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// storeCategoryID := guuid.New()
	// if _, err := stmt.Exec(storeID); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit()
}

//GetStoreByID
func (s *StoreRepoImpl) GetStoreByID(id string) (*models.Store, error) {
	stmt, err := s.db.Prepare(utils.SELECT_STORE_BY_ID)
	store := models.Store{}
	if err != nil {
		return &store, err
	}
	errQuery := stmt.QueryRow(id).Scan(&store.StoreID, &store.StoreName, &store.StoreAddress, &store.StoreOwner, &store.StoreUsername, &store.StoreStatus, &store.StoreImage, &store.StoreCategory.StoreCategoryID, &store.StoreCategory.StoreCategoryName)

	if errQuery != nil {
		return &store, err
	}

	defer stmt.Close()
	return &store, nil
}

// //GetAllUser for admin
func (s *StoreRepoImpl) GetAllStore() ([]*models.Store, error) {
	stmt, err := s.db.Prepare(utils.SELECT_ALL_STORE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	listStore := []*models.Store{}
	for rows.Next() {
		store := models.Store{}
		err := rows.Scan(&store.StoreID, &store.StoreName, &store.StoreAddress, &store.StoreOwner, &store.StoreStatus, &store.StoreUsername, &store.StorePassword, &store.StoreImage, &store.StoreCategory.StoreCategoryID, &store.StoreCategory.StoreCategoryName)
		if err != nil {
			println(err)
			return nil, err
		}
		listStore = append(listStore, &store)
	}
	return listStore, nil
}

// // Update Store for profil
func (s *StoreRepoImpl) UpdateStore(id string, store *models.Store) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_STORE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(store.StoreName, store.StoreCategory.StoreCategoryID, store.StoreAddress, store.StoreOwner, store.StoreUsername, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// // Delete User for admin
func (s *StoreRepoImpl) DeleteStore(storeID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_STORE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(storeID)
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

// // handle login / auth
// func (u *UserRepoImpl) Auth(username, password string) (*models.Auth, error) {
// 	stmt, err := u.db.Prepare(utils.LOGIN)
// 	user := models.Auth{}
// 	if err != nil {
// 		return &user, err
// 	}
// 	errQuery := stmt.QueryRow(username, password).Scan(user.UserID, user.UserLevelID, user.UserStatus)
// 	if errQuery != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	return &user, nil
// }

// //login 2
func (s *StoreRepoImpl) Auth(username string) (*models.Store, error) {
	fmt.Println("MASUK REPO", username)
	stmt, err := s.db.Prepare(utils.STORE_AUTH)
	store := models.Store{}
	if err != nil {
		fmt.Println(err)
		return &store, err
	}
	errQuery := stmt.QueryRow(username).Scan(&store.StoreID, &store.StoreUsername, &store.StoreCategory.StoreCategoryID, &store.StoreAddress, &store.StoreOwner, &store.StoreStatus, &store.StoreUsername, &store.StorePassword, &store.StoreImage)
	log.Println(errQuery)
	if errQuery != nil {
		return &store, err
	}
	fmt.Println(store.StorePassword)
	defer stmt.Close()
	return &store, nil
}
