package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type UsersUseCaseImpl struct {
	usersRepo repositories.UsersRepo
}

func (u UsersUseCaseImpl) Auth(userModels *models.UserModels) ([]*models.Auth, error) {
	//err := validation.CheckEmpty(userModels.UserName, userModels.UserPassword)
	//if err != nil {
	//	return nil, err
	//}
	auth, err := u.usersRepo.Auth(userModels)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func InitUsersUseCase(users repositories.UsersRepo) UsersUseCases {
	return &UsersUseCaseImpl{users}
}
