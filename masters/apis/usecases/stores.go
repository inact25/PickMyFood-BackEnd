package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type StoreUsecaseImpl struct {
	storeRepo repositories.StoresRepo
}

func (s StoreUsecaseImpl) GetStores() ([]*models.StoreModels, error) {
	stores, err := s.storeRepo.GetStores()
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (s StoreUsecaseImpl) GetStoreByID(ID string) (*models.StoreModels, error) {
	stores, err := s.storeRepo.GetStoreByID(ID)

	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (s StoreUsecaseImpl) PostStore(d models.StoreModels) (*models.StoreModels, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.storeRepo.PostStore(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s StoreUsecaseImpl) UpdateStore(ID string, data models.StoreModels) (*models.StoreModels, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.storeRepo.UpdateStore(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s StoreUsecaseImpl) DeleteStore(ID string) (*models.StoreModels, error) {
	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	_, err := s.storeRepo.GetStoreByID(ID)
	if err != nil {
		return nil, err
	}

	result, err := s.storeRepo.DeleteStore(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s StoreUsecaseImpl) GetStoresCategory() ([]*models.StoreCategory, error) {
	storesCategory, err := s.storeRepo.GetStoresCategory()
	if err != nil {
		return nil, err
	}
	return storesCategory, nil
}

func InitStoreUsecase(storeRepo repositories.StoresRepo) StoresUseCases {
	return &StoreUsecaseImpl{storeRepo}
}
