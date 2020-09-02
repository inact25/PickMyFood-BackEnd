package repositories

import (
	"database/sql"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
)

type PoinRepoImpl struct {
	db *sql.DB
}

func (s *PoinRepoImpl) GetPoints() ([]*models.PoinModels, error) {
	var points []*models.PoinModels
	query := "SELECT * FROM tb_poin"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		point := models.PoinModels{}
		err := rows.Scan(&point.PoinID, &point.StoreID)

		if err != nil {
			return nil, err
		}

		points = append(points, &point)

	}

	return points, nil
}

func (s *PoinRepoImpl) GetPointByID(ID string) ([]*models.PoinModels, error) {
	var points []*models.PoinModels
	query := "SELECT * FROM tb_poin WHERE poin_id = ?"
	rows, err := s.db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		point := models.PoinModels{}
		err := rows.Scan(&point.PoinID, &point.StoreID)

		if err != nil {
			return nil, err
		}

		points = append(points, &point)

	}

	return points, nil
}

func InitPoinRepoImpl(db *sql.DB) PoinRepo {
	return &PoinRepoImpl{db}

}
