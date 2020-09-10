package storeCategoryUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	storeCategoryRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/storeCategory"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type StoreCategoryUsecaseImpl struct {
	storeCategoryRepo storeCategoryRepositories.StoreCategoryRepo
}

func InitStoreCategoryUseCase(storeCategory storeCategoryRepositories.StoreCategoryRepo) StoreCategoryUsecase {
	return &StoreCategoryUsecaseImpl{storeCategory}
}

// AddStore usecase
func (sc *StoreCategoryUsecaseImpl) AddStoreCategory(storeCategory *models.StoreCategory) error {
	err := validation.CheckEmpty(storeCategory)
	if err != nil {
		return err
	}
	error := sc.storeCategoryRepo.AddStoreCategory(storeCategory)
	if error != nil {
		return error
	}
	return nil
}

// GetUserById
func (sc *StoreCategoryUsecaseImpl) GetStoreCategoryByID(storeCategoryID string) (*models.StoreCategory, error) {
	storeCategory, err := sc.storeCategoryRepo.GetStoreCategoryByID(storeCategoryID)
	if err != nil {
		return nil, err
	}
	return storeCategory, nil
}

func (sc *StoreCategoryUsecaseImpl) GetAllStoreCategory() ([]*models.StoreCategory, error) {
	listStoreCategory, err := sc.storeCategoryRepo.GetAllStoreCategory()
	if err != nil {
		return nil, err
	}
	return listStoreCategory, nil
}

func (sc *StoreCategoryUsecaseImpl) UpdateStoreCategory(id string, storeCategory *models.StoreCategory) error {
	err := validation.CheckEmpty(storeCategory)
	if err != nil {
		return err
	}
	error := sc.storeCategoryRepo.UpdateStoreCategory(id, storeCategory)
	if error != nil {
		return error
	}
	return nil
}

func (sc *StoreCategoryUsecaseImpl) DeleteStoreCategory(storeCategoryID string) error {
	err := sc.storeCategoryRepo.DeleteStoreCategory(storeCategoryID)
	if err != nil {
		return err
	}
	return nil
}
