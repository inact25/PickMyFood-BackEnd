package userRepositories

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type UserRepo interface {
	AddUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetAllUser(keyword, page, limit string) ([]*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
	Auth(username, password string) (*models.Auth, error)
	ReadUserByUsername(username string) (*models.User, error)
	UserNonAktif(keyword, page, limit string) ([]*models.User, error)
	ChangeActive(userID string) error
	ChangeProfile(id string, user *models.User) error
}
