package feedbackRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type FeedbackRepo interface {
	GetFeedbacks() ([]*models.FeedbackModels, error)
	GetFeedbackByID(ID string) (*models.FeedbackModels, error)
	PostFeedback(d *models.FeedbackModels) error
	UpdateFeedback(ID string, data *models.FeedbackModels) error
	DeleteFeedback(ID string) error
}
