package repositories

import (
	"database/sql"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
)

type RatingRepoImpl struct {
	db *sql.DB
}

func (s *RatingRepoImpl) GetRatings() ([]*models.RatingModels, error) {
	var ratings []*models.RatingModels
	query := "SELECT * FROM tb_rating"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rating := models.RatingModels{}
		err := rows.Scan(&rating.RatingID, &rating.StoreID, &rating.UserID, &rating.RatingValue, &rating.RatingDescription, &rating.RatingCreated)

		if err != nil {
			return nil, err
		}

		ratings = append(ratings, &rating)

	}

	return ratings, nil
}

func (s *RatingRepoImpl) GetRatingByID(ID string) ([]*models.RatingModels, error) {
	var ratings []*models.RatingModels
	query := "SELECT * FROM tb_rating WHERE rating_id = ?"
	rows, err := s.db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rating := models.RatingModels{}
		err := rows.Scan(&rating.RatingID, &rating.StoreID, &rating.UserID, &rating.RatingValue, &rating.RatingDescription, &rating.RatingCreated)

		if err != nil {
			return nil, err
		}

		ratings = append(ratings, &rating)

	}

	return ratings, nil
}

func InitRatingRepoImpl(db *sql.DB) RatingRepo {
	return &RatingRepoImpl{db}

}
