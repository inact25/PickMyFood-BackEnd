package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

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

func (s *RatingRepoImpl) GetRatingByID(ID string) (*models.RatingModels, error) {
	results := s.db.QueryRow("SELECT * FROM tb_rating WHERE rating_id = ?", ID)

	var d models.RatingModels
	err := results.Scan(&d.RatingID, &d.StoreID, &d.UserID, &d.RatingValue, &d.RatingDescription, &d.RatingDescription)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *RatingRepoImpl) PostRating(d models.RatingModels) (*models.RatingModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO tb_rating(rating_id, store_id, user_id, rating_value, rating_description, rating_created) VALUES (?, ?, ?, ?, ?, ?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.RatingID, d.StoreID, d.UserID, d.RatingValue, d.RatingDescription, d.RatingCreated)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return s.GetRatingByID(strconv.Itoa(int(lastInsertID)))
}

func (s *RatingRepoImpl) UpdateRating(ID string, data models.RatingModels) (*models.RatingModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE tb_rating SET store_id=?, user_id=?, rating_value=?, rating_description=?, rating_created=? WHERE rating_id=?`,
		data.StoreID, data.UserID, data.RatingValue, data.RatingDescription, data.RatingCreated, ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.GetRatingByID(ID)
}

func (s *RatingRepoImpl) DeleteRating(ID string) (*models.RatingModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM tb_rating WHERE rating_id = ?", ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return s.GetRatingByID(ID)

}

func InitRatingRepoImpl(db *sql.DB) RatingRepo {
	return &RatingRepoImpl{db}

}
