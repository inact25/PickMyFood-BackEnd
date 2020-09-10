package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type PoinUseCases interface {
	GetPoints() ([]*models.PoinModels, error)
	GetPointByID(ID string) (*models.PoinModels, error)
	PostPoint(d *models.PoinModels) error
	UpdatePoint(data *models.PoinModels, ID string) error
	DeletePoint(ID string) error

	UpdateUserPoint(ID string, data *models.User) error
}
