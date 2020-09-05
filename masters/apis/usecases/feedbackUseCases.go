package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type FeedbackUseCases interface {
	GetFeedbacks() ([]*models.FeedbackModels, error)
	GetFeedbackByID(ID string) (*models.FeedbackModels, error)
	PostFeedback(d models.FeedbackModels) (*models.FeedbackModels, error)
	UpdateFeedback(ID string, data models.FeedbackModels) (*models.FeedbackModels, error)
	DeleteFeedback(ID string) (*models.FeedbackModels, error)
}
