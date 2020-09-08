package userUsecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type UserUseCase interface {
	// Auth(*models.UserModels) ([]*models.Auth, error)
	AddUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetAllUser(keyword, page, limit string) ([]*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(userID string) error
	Auth(username, password string) (*models.Auth, error)
	ReadUserByUsername(username string) (*models.Auth, error)
}
