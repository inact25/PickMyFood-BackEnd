package repositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type RatingRepo interface {
	GetRatings() ([]*models.RatingModels, error)
	GetRatingByID(ID string) ([]*models.RatingModels, error)
}
