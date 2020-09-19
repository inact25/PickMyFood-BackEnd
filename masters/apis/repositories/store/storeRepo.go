package storerepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type StoreRepo interface {
	AddStore(store *models.Store) error
	GetStoreByID(id string) (*models.Store, error)
	GetAllStore(keyword string) ([]*models.Store, error)
	UpdateStore(id string, Store *models.Store) error
	DeleteStore(id string) error
	Auth(username string) (*models.Store, error)
	GetStoreNonAktif() ([]*models.Store, error)
	ChangeActive(id string) error
}
