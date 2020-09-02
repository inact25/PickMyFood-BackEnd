package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type RatingUseCases interface {
	GetRatings() ([]*models.RatingModels, error)
	GetRatingByID(ID string) ([]*models.RatingModels, error)
}
