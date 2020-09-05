package storeCategoryRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type StoreCategoryRepo interface {
	AddStoreCategory(storeCategory *models.StoreCategory) error
	GetStoreCategoryByID(id string) (*models.StoreCategory, error)
	GetAllStoreCategory() ([]*models.StoreCategory, error)
	UpdateStoreCategory(id string, storeCategory *models.StoreCategory) error
	DeleteStoreCategory(id string) error
}
