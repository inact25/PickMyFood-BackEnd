package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
)

type RatingUsecaseImpl struct {
	ratingRepo repositories.RatingRepo
}

func (s RatingUsecaseImpl) GetRatings() ([]*models.RatingModels, error) {
	ratings, err := s.ratingRepo.GetRatings()
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (s RatingUsecaseImpl) GetRatingByID(ID string) ([]*models.RatingModels, error) {
	ratings, err := s.ratingRepo.GetRatingByID(ID)

	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func InitRatingUsecase(ratingRepo repositories.RatingRepo) RatingUseCases {
	return &RatingUsecaseImpl{ratingRepo}
}
