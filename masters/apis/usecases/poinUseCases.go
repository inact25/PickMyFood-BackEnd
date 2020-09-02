package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type PoinUseCases interface {
	GetPoints() ([]*models.PoinModels, error)
	GetPointByID(ID string) ([]*models.PoinModels, error)
}
