package repositories

import (
	"database/sql"
	"errors"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type PoinRepoImpl struct {
	db *sql.DB
}

func (s *PoinRepoImpl) GetPoints() ([]*models.PoinModels, error) {
	var points []*models.PoinModels
	query := utils.GET_ALL_POINT
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

func (s *PoinRepoImpl) GetPointByID(ID string) (*models.PoinModels, error) {
	results := s.db.QueryRow(utils.GET_POINT_BY_ID, ID)

	var d models.PoinModels
	err := results.Scan(&d.PoinID, &d.StoreID)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *PoinRepoImpl) PostPoint(d *models.PoinModels, ID string) error {
	pointID := guuid.New()
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.POST_POINT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(pointID, d.StoreID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *PoinRepoImpl) UpdatePoint(data *models.PoinModels, ID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_POINT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(data.StoreID, ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	println("UPDATE POINT")

	return tx.Commit()
}

func (s *PoinRepoImpl) DeletePoint(ID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_POINT)
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

func (s *PoinRepoImpl) UpdateUserPoint(ID string, data *models.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_USER_POINT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	println("UPDATE USER POINT")

	return tx.Commit()
}

func InitPoinRepoImpl(db *sql.DB) PoinRepo {
	return &PoinRepoImpl{db}
}
