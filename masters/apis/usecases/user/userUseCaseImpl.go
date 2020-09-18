package userUsecases

import (
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	userRepositories "github.com/inact25/PickMyFood-BackEnd/masters/apis/repositories/user"
	"github.com/inact25/PickMyFood-BackEnd/utils/validation"
)

type UserUseCaseImpl struct {
	userRepo userRepositories.UserRepo
}

func InitUsersUseCase(users userRepositories.UserRepo) UserUseCase {
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

func (u *UserUseCaseImpl) GetAllUser(keyword, page, limit string) ([]*models.User, error) {
	listUser, err := u.userRepo.GetAllUser(keyword, page, limit)
	if err != nil {
		return nil, err
	}
	return listUser, nil
}

func (u *UserUseCaseImpl) UpdateUser(id string, user *models.User) error {
	err := validation.CheckEmpty(user.UserAddress, user.UserLastName, user.UserAddress, user.UserPhone, user.UserImage, user.UserStatus, user.Auth.Username, user.Auth.Password)
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
	auth, err := u.userRepo.Auth(username, password)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
func (u *UserUseCaseImpl) ReadUserByUsername(username string) (*models.User, error) {
	user, err := u.userRepo.ReadUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *UserUseCaseImpl) UserNonAktif(keyword, page, limit string) ([]*models.User, error) {
	listUser, err := u.userRepo.UserNonAktif(keyword, page, limit)
	if err != nil {
		return nil, err
	}
	return listUser, nil
}
func (u *UserUseCaseImpl) ChangeActive(userID string) error {
	err := u.userRepo.ChangeActive(userID)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUseCaseImpl) ChangeProfile(id string, user *models.User) error {
	err := validation.CheckEmpty(user)
	if err != nil {
		return err
	}
	error := u.userRepo.ChangeProfile(id, user)
	if error != nil {
		return error
	}
	return nil
}
