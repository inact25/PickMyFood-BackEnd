package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type FeedbackUseCases interface {
	GetFeedbacks() ([]*models.FeedbackModels, error)
	GetFeedbackByID(ID string) (*models.FeedbackModels, error)
	PostFeedback(d *models.FeedbackModels, ID string) error
	UpdateFeedback(data *models.FeedbackModels, ID string) error
	DeleteFeedback(ID string) error
}
