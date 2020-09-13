package ratingRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type RatingRepo interface {
	GetRatings() ([]*models.RatingModels, error)
	GetRatingByID(ID string) (*models.RatingModels, error)
	PostRating(d *models.RatingModels) error
	UpdateRating(data *models.RatingModels, ID string) error
	DeleteRating(ID string) error
}
