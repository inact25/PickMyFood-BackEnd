package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
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

func (s RatingUsecaseImpl) GetRatingByID(ID string) (*models.RatingModels, error) {
	ratings, err := s.ratingRepo.GetRatingByID(ID)

	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (s RatingUsecaseImpl) PostRating(d *models.RatingModels, ID string) error {
	err := validation.CheckEmpty(d)
	if err != nil {
		return err
	}
	error := s.ratingRepo.PostRating(d, ID)
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

func InitRatingUsecase(ratingRepo repositories.RatingRepo) RatingUseCases {
	return &RatingUsecaseImpl{ratingRepo}
}
