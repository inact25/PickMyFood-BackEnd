package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type UsersUseCases interface {
	Auth(*models.UserModels) ([]*models.Auth, error)
}
