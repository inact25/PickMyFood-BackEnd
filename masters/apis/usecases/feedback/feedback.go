package feedbackUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	feedbackRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/feedback"
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type FeedbackUsecaseImpl struct {
	feedbackRepo feedbackRepositories.FeedbackRepo
}

func (s FeedbackUsecaseImpl) GetFeedbacks() ([]*models.FeedbackModels, error) {
	feedbacks, err := s.feedbackRepo.GetFeedbacks()
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (s FeedbackUsecaseImpl) GetFeedbackByID(ID string) (*models.FeedbackModels, error) {
	feedbacks, err := s.feedbackRepo.GetFeedbackByID(ID)

	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (s FeedbackUsecaseImpl) PostFeedback(d *models.FeedbackModels) error {
	d.FeedbackCreated = utils.GetTimeNow()
	err := validation.CheckEmpty(d)
	if err != nil {
		return err
	}
	error := s.feedbackRepo.PostFeedback(d)
	if error != nil {
		return error
	}
	return nil
}

func (s FeedbackUsecaseImpl) UpdateFeedback(data *models.FeedbackModels, ID string) error {
	if err := validator.Validate(data); err != nil {
		return err
	}

	err := s.feedbackRepo.UpdateFeedback(ID, data)
	if err != nil {
		return err
	}
	return nil
}

func (s FeedbackUsecaseImpl) DeleteFeedback(ID string) error {
	err := s.feedbackRepo.DeleteFeedback(ID)
	if err != nil {
		return err
	}
	return nil
}

func InitFeedbackUsecase(feedbackRepo feedbackRepositories.FeedbackRepo) FeedbackUseCases {
	return &FeedbackUsecaseImpl{feedbackRepo}
}
