package repositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type UsersRepo interface {
	Auth(*models.UserModels) ([]*models.Auth, error)
}
