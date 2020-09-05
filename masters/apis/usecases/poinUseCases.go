package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type PoinUseCases interface {
	GetPoints() ([]*models.PoinModels, error)
	GetPointByID(ID string) (*models.PoinModels, error)
	PostPoint(d models.PoinModels) (*models.PoinModels, error)
	UpdatePoint(ID string, data models.PoinModels) (*models.PoinModels, error)
	DeletePoint(ID string) (*models.PoinModels, error)
}
