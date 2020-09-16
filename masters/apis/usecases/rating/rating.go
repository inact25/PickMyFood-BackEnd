package ratingUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	ratingRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/rating"
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type RatingUsecaseImpl struct {
	ratingRepo ratingRepositories.RatingRepo
}

func (s RatingUsecaseImpl) GetRatings(storeID string) ([]*models.RatingModels, error) {
	ratings, err := s.ratingRepo.GetRatings(storeID)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (s RatingUsecaseImpl) GetRatingByID(ID string) (*models.RatingModels, error) {
	ratings, err := s.ratingRepo.GetRatingByID(ID)

	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (s RatingUsecaseImpl) PostRating(d *models.RatingModels) error {
	d.RatingCreated = utils.GetTimeNow()
	err := validation.CheckEmpty(d)
	if err != nil {
		return err
	}
	error := s.ratingRepo.PostRating(d)
	if error != nil {
		return error
	}
	return nil
}

func (s RatingUsecaseImpl) UpdateRating(data *models.RatingModels, ID string) error {
	if err := validator.Validate(data); err != nil {
		return err
	}

	err := s.ratingRepo.UpdateRating(data, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s RatingUsecaseImpl) DeleteRating(ID string) error {
	err := s.ratingRepo.DeleteRating(ID)
	if err != nil {
		return err
	}
	return nil
}

func InitRatingUsecase(ratingRepo ratingRepositories.RatingRepo) RatingUseCases {
	return &RatingUsecaseImpl{ratingRepo}
}
