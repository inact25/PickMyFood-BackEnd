package repositories

import (
	"database/sql"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u UserRepoImpl) Auth(userModels *models.UserModels) ([]*models.Auth, error) {
	panic("implement me")
}

func InitUserRepoImpl(db *sql.DB) UsersRepo {
	return &UserRepoImpl{db}

}
