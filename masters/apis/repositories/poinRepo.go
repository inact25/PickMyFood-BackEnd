package repositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type PoinRepo interface {
	GetPoints() ([]*models.PoinModels, error)
	GetPointByID(ID string) ([]*models.PoinModels, error)
}
