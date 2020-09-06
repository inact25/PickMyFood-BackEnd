package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type PoinRepoImpl struct {
	db *sql.DB
}

func (s *PoinRepoImpl) GetPoints() ([]*models.PoinModels, error) {
	var points []*models.PoinModels
	query := queryConstant.GET_ALL_POINT
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
	results := s.db.QueryRow(queryConstant.GET_POINT_BY_ID, ID)

	var d models.PoinModels
	err := results.Scan(&d.PoinID, &d.StoreID)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *PoinRepoImpl) PostPoint(d models.PoinModels) (*models.PoinModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(queryConstant.POST_POINT)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.PoinID, d.StoreID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return s.GetPointByID(strconv.Itoa(int(lastInsertID)))
}

func (s *PoinRepoImpl) UpdatePoint(ID string, data models.PoinModels) (*models.PoinModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(queryConstant.UPDATE_POINT,
		data.StoreID, ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.GetPointByID(ID)
}

func (s *PoinRepoImpl) DeletePoint(ID string) (*models.PoinModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(queryConstant.DELETE_POINT, ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return s.GetPointByID(ID)

}

func InitPoinRepoImpl(db *sql.DB) PoinRepo {
	return &PoinRepoImpl{db}

}
