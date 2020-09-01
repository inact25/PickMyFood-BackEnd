package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
)

type StoreUsecaseImpl struct {
	storeRepo repositories.StoresRepository
}

func (s StoreUsecaseImpl) GetStores() ([]*models.StoreModels, error) {
	stores, err := s.storeRepo.GetStores()
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (s StoreUsecaseImpl) GetStoreByID(ID string) ([]*models.StoreModels, error) {
	stores, err := s.storeRepo.GetStoreByID(ID)

	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (s StoreUsecaseImpl) DeleteStore(ID string) ([]*models.StoreModels, error) {
	stores, err := s.storeRepo.DeleteStore(ID)

	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (s StoreUsecaseImpl) GetStoresCategory() ([]*models.StoreCategory, error) {
	storesCategory, err := s.storeRepo.GetStoresCategory()
	if err != nil {
		return nil, err
	}
	return storesCategory, nil
}

func InitStoreUsecase(storeRepo repositories.StoresRepository) StoresUseCases {
	return &StoreUsecaseImpl{storeRepo}
}
