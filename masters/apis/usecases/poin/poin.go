package poinUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	poinRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/poin"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type PoinUsecaseImpl struct {
	poinRepo poinRepositories.PoinRepo
}

func (s PoinUsecaseImpl) GetPoints() ([]*models.PoinModels, error) {
	points, err := s.poinRepo.GetPoints()
	if err != nil {
		return nil, err
	}
	return points, nil
}

func (s PoinUsecaseImpl) GetPointByID(ID string) (*models.PoinModels, error) {
	points, err := s.poinRepo.GetPointByID(ID)

	if err != nil {
		return nil, err
	}
	return points, nil
}

func (s PoinUsecaseImpl) PostPoint(d *models.PoinModels) error {
	err := validation.CheckEmpty(d)
	if err != nil {
		return err
	}
	error := s.poinRepo.PostPoint(d)
	if error != nil {
		return error
	}
	return nil
}

func (s PoinUsecaseImpl) UpdatePoint(data *models.PoinModels, ID string) error {
	if err := validator.Validate(data); err != nil {
		return err
	}

	err := s.poinRepo.UpdatePoint(data, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s PoinUsecaseImpl) DeletePoint(ID string) error {
	err := s.poinRepo.DeletePoint(ID)
	if err != nil {
		return err
	}
	return nil
}

func (s PoinUsecaseImpl) UpdateUserPoint(ID string, data *models.User) error {
	if err := validator.Validate(data); err != nil {
		return err
	}

	err := s.poinRepo.UpdateUserPoint(ID, data)
	if err != nil {
		return err
	}
	return nil
}

func InitPoinUsecase(poinRepo poinRepositories.PoinRepo) PoinUseCases {
	return &PoinUsecaseImpl{poinRepo}
}
