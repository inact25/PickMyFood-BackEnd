package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
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

func (s PoinUsecaseImpl) GetPointByID(ID string) ([]*models.PoinModels, error) {
	points, err := s.poinRepo.GetPointByID(ID)

	if err != nil {
		return nil, err
	}
	return points, nil
}

func InitPoinUsecase(poinRepo repositories.PoinRepo) PoinUseCases {
	return &PoinUsecaseImpl{poinRepo}
}
