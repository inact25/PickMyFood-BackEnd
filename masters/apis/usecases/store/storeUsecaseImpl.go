package storeusecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	storerepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/store"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type StoreUsecaseImpl struct {
	storeRepo storerepositories.StoreRepo
}

func InitStoreUseCase(store storerepositories.StoreRepo) StoreUsecase {
	return &StoreUsecaseImpl{store}
}

// AddStore usecase
func (s *StoreUsecaseImpl) AddStore(store *models.Store) error {
	err := validation.CheckEmpty(store)
	if err != nil {
		return err
	}
	error := s.storeRepo.AddStore(store)
	if error != nil {
		return error
	}
	return nil
}

// GetUserById
func (s *StoreUsecaseImpl) GetStoreByID(storeID string) (*models.Store, error) {
	store, err := s.storeRepo.GetStoreByID(storeID)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *StoreUsecaseImpl) GetAllStore() ([]*models.Store, error) {
	listStore, err := s.storeRepo.GetAllStore()
	if err != nil {
		return nil, err
	}
	return listStore, nil
}

func (s *StoreUsecaseImpl) UpdateStore(id string, store *models.Store) error {
	err := validation.CheckEmpty(store)
	if err != nil {
		return err
	}
	error := s.storeRepo.UpdateStore(id, store)
	if error != nil {
		return error
	}
	return nil
}

func (s *StoreUsecaseImpl) DeleteStore(storeID string) error {
	err := s.storeRepo.DeleteStore(storeID)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreUsecaseImpl) Auth(username string) (*models.Store, error) {
	//err := validation.CheckEmpty(userModels.UserName, userModels.UserPassword)
	//if err != nil {
	//	return nil, err
	//}
	auth, err := s.storeRepo.Auth(username)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
