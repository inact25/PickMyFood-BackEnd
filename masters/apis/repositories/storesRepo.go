package repositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type StoresRepo interface {
	GetStores() ([]*models.StoreModels, error)
	GetStoreByID(ID string) ([]*models.StoreModels, error)
	DeleteStore(ID string) ([]*models.StoreModels, error)

	GetStoresCategory() ([]*models.StoreCategory, error)
}
