package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type PoinUsecaseImpl struct {
	poinRepo repositories.PoinRepo
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

func (s PoinUsecaseImpl) PostPoint(d models.PoinModels) (*models.PoinModels, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.poinRepo.PostPoint(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s PoinUsecaseImpl) UpdatePoint(ID string, data models.PoinModels) (*models.PoinModels, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.poinRepo.UpdatePoint(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s PoinUsecaseImpl) DeletePoint(ID string) (*models.PoinModels, error) {
	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	_, err := s.poinRepo.GetPointByID(ID)
	if err != nil {
		return nil, err
	}

	result, err := s.poinRepo.DeletePoint(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func InitPoinUsecase(poinRepo repositories.PoinRepo) PoinUseCases {
	return &PoinUsecaseImpl{poinRepo}
}
