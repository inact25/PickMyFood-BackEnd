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

func (s RatingUsecaseImpl) PostRating(d models.RatingModels) (*models.RatingModels, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.ratingRepo.PostRating(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s RatingUsecaseImpl) UpdateRating(ID string, data models.RatingModels) (*models.RatingModels, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.ratingRepo.UpdateRating(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s RatingUsecaseImpl) DeleteRating(ID string) (*models.RatingModels, error) {
	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	_, err := s.ratingRepo.GetRatingByID(ID)
	if err != nil {
		return nil, err
	}

	result, err := s.ratingRepo.DeleteRating(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func InitRatingUsecase(ratingRepo repositories.RatingRepo) RatingUseCases {
	return &RatingUsecaseImpl{ratingRepo}
}
