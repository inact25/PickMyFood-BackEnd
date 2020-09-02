package usecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type UserUseCaseImpl struct {
	userRepo repositories.UserRepo
}

func InitUsersUseCase(users repositories.UserRepo) UserUseCase {
	return &UserUseCaseImpl{users}
}

// AddUser usecase
func (u *UserUseCaseImpl) AddUser(user *models.User) error {
	err := validation.CheckEmpty(user)
	if err != nil {
		return err
	}
	error := u.userRepo.AddUser(user)
	if error != nil {
		return error
	}
	return nil
}

// GetUserById
func (u *UserUseCaseImpl) GetUserByID(userID string) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCaseImpl) GetAllUser() ([]*models.User, error) {
	listUser, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return listUser, nil
}

func (u *UserUseCaseImpl) UpdateUser(id string, user *models.User) error {
	err := validation.CheckEmpty(user)
	if err != nil {
		return err
	}
	error := u.userRepo.UpdateUser(id, user)
	if error != nil {
		return error
	}
	return nil
}

func (u *UserUseCaseImpl) DeleteUser(userID string) error {
	err := u.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCaseImpl) Auth(username, password string) (*models.Auth, error) {
	//err := validation.CheckEmpty(userModels.UserName, userModels.UserPassword)
	//if err != nil {
	//	return nil, err
	//}
	auth, err := u.userRepo.Auth(username, password)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
func (u *UserUseCaseImpl) ReadUserByUsername(username string) (*models.Auth, error) {
	user, err := u.userRepo.ReadUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
