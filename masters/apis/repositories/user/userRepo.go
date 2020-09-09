package userRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type UserRepo interface {
	AddUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetAllUser() ([]*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
	Auth(username, password string) (*models.Auth, error)
	ReadUserByUsername(username string) (*models.Auth, error)
}
