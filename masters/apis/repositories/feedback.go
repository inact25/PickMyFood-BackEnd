package repositories

import (
	"database/sql"
	"errors"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type FeedbackRepoImpl struct {
	db *sql.DB
}

func (s *FeedbackRepoImpl) GetFeedbacks() ([]*models.FeedbackModels, error) {
	var feedbacks []*models.FeedbackModels
	query := utils.GET_ALL_FEEDBACK
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		feedback := models.FeedbackModels{}
		err := rows.Scan(&feedback.FeedbackID, &feedback.StoreID, &feedback.UserID, &feedback.FeedbackValue, &feedback.FeedbackCreated)

		if err != nil {
			return nil, err
		}

		feedbacks = append(feedbacks, &feedback)

	}

	return feedbacks, nil
}

func (s *FeedbackRepoImpl) GetFeedbackByID(ID string) (*models.FeedbackModels, error) {
	results := s.db.QueryRow(utils.GET_FEEDBACK_BY_ID, ID)

	var d models.FeedbackModels
	err := results.Scan(&d.FeedbackID, &d.StoreID, &d.UserID, &d.FeedbackValue, &d.FeedbackCreated)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *FeedbackRepoImpl) PostFeedback(d *models.FeedbackModels, ID string) error {
	feedbackID := guuid.New()
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.POST_FEEDBACK)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(feedbackID, d.StoreID, d.UserID, d.FeedbackValue, d.FeedbackCreated); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *FeedbackRepoImpl) UpdateFeedback(ID string, data *models.FeedbackModels) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_FEEDBACK)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(data.StoreID, data.UserID, data.FeedbackValue, data.FeedbackCreated, ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	println("UPDATE FEEDBACK")

	return tx.Commit()
}

func (s *FeedbackRepoImpl) DeleteFeedback(ID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_FEEDBACK)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, id tidak di temukan")
	}

	return tx.Commit()
}

func InitFeedbackImpl(db *sql.DB) FeedbackRepo {
	return &FeedbackRepoImpl{db}

}
