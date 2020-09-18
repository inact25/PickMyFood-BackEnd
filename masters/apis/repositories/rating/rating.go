package ratingRepositories

import (
	"database/sql"
	"errors"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type RatingRepoImpl struct {
	db *sql.DB
}

func (s *RatingRepoImpl) GetRatings(storeID string) ([]*models.RatingModels, error) {
	var ratings []*models.RatingModels
	query := utils.GET_ALL_RATING
	rows, err := s.db.Query(query, storeID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rating := models.RatingModels{}
		err := rows.Scan(&rating.RatingID, &rating.StoreID, &rating.UserID, &rating.RatingValue, &rating.RatingDescription, &rating.RatingCreated, &rating.UserFirstname, &rating.UserLastname)

		if err != nil {
			return nil, err
		}

		ratings = append(ratings, &rating)

	}

	return ratings, nil
}

func (s *RatingRepoImpl) GetRatingByID(ID string) (*models.RatingModels, error) {
	results := s.db.QueryRow(utils.GET_RATING_BY_ID, ID)

	var d models.RatingModels
	err := results.Scan(&d.RatingID, &d.StoreID, &d.UserID, &d.RatingValue, &d.RatingDescription, &d.RatingCreated)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *RatingRepoImpl) PostRating(d *models.RatingModels) error {
	ratingID := guuid.New()
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.POST_RATING)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(ratingID, d.StoreID, d.UserID, d.RatingValue, d.RatingDescription, d.RatingCreated); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *RatingRepoImpl) UpdateRating(data *models.RatingModels, ID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_RATING)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(data.StoreID, data.UserID, data.RatingValue, data.RatingDescription, data.RatingCreated, ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *RatingRepoImpl) DeleteRating(ID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_RATING)
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

func InitRatingRepoImpl(db *sql.DB) RatingRepo {
	return &RatingRepoImpl{db}

}
