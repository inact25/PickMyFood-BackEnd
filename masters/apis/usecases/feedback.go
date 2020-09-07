package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"

	"gopkg.in/validator.v2"
)

type FeedbackUsecaseImpl struct {
	feedbackRepo repositories.FeedbackRepo
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

func (s FeedbackUsecaseImpl) PostFeedback(d models.FeedbackModels) (*models.FeedbackModels, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.feedbackRepo.PostFeedback(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s FeedbackUsecaseImpl) UpdateFeedback(ID string, data models.FeedbackModels) (*models.FeedbackModels, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(ID); err != nil {
		return nil, err
	}

	result, err := s.feedbackRepo.UpdateFeedback(ID, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s FeedbackUsecaseImpl) DeleteFeedback(ID string) error {
	err := s.feedbackRepo.DeleteFeedback(ID)
	if err != nil {
		return err
	}
	return nil
}

func InitFeedbackUsecase(feedbackRepo repositories.FeedbackRepo) FeedbackUseCases {
	return &FeedbackUsecaseImpl{feedbackRepo}
}
