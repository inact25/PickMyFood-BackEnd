package userRepositories

import (
	"database/sql"
	"fmt"
	"log"

	"errors"

	guuid "github.com/google/uuid"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	utils "github.com/inact25/PickMyFood-BackEnd/utils/queryConstant"
)

type UserRepoImpl struct {
	db *sql.DB
}

func InitUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{db: db}
}

func (u *UserRepoImpl) AddUser(user *models.User) error {
	id := guuid.New()
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(id, user.UserFirstName, user.UserLastName, user.UserAddress, user.UserPhone, user.UserImage, user.UserStatus); err != nil {
		tx.Rollback()
		return err
	}

	stmt, err = tx.Prepare(utils.INSERT_AUTH)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	// profileID := guuid.New()
	if _, err := stmt.Exec(user.Auth.Username, user.Auth.Password, id); err != nil {
		tx.Rollback()
		return err
	}

	stmt, err = tx.Prepare(utils.INSERT_WALLET)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	walletID := guuid.New()
	if _, err := stmt.Exec(walletID, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// GetUserByID for view profil user
func (u *UserRepoImpl) GetUserByID(id string) (*models.User, error) {
	stmt, err := u.db.Prepare(utils.SELECT_USER_BY_ID)
	user := models.User{}
	if err != nil {
		return &user, err
	}
	errQuery := stmt.QueryRow(id).Scan(&user.Auth.Username, &user.Auth.Password, &user.UserImage, &user.UserPoin, &user.UserStatus, &user.UserFirstName, &user.UserLastName, &user.UserPhone, &user.UserAddress)

	if errQuery != nil {
		return &user, err
	}

	defer stmt.Close()
	return &user, nil
}

//GetAllUser for admin
func (u *UserRepoImpl) GetAllUser() ([]*models.User, error) {
	stmt, err := u.db.Prepare(utils.SELECT_ALL_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	listUser := []*models.User{}
	for rows.Next() {
		p := models.User{}
		err := rows.Scan(&p.UserID, &p.UserFirstName, &p.UserLastName, &p.UserAddress, &p.UserPhone, &p.UserPoin, &p.UserImage, &p.UserStatus, &p.Auth.Username, &p.Auth.Password, &p.Auth.UserLevelID, &p.Auth.UserStatus)
		if err != nil {
			return nil, err
		}
		listUser = append(listUser, &p)
	}
	return listUser, nil
}

// Update User for profil
func (u *UserRepoImpl) UpdateUser(id string, user *models.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(user.UserFirstName, user.UserLastName, user.UserAddress, user.UserPhone, user.UserImage, user.UserStatus, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt, err = tx.Prepare(utils.UPDATE_AUTH)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(user.Auth.Username, user.Auth.Password, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// Delete User for admin
func (u *UserRepoImpl) DeleteUser(userID string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_AUTH)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, user id tidak di temukan")
	}

	return tx.Commit()
}

// handle login / auth
func (u *UserRepoImpl) Auth(username, password string) (*models.Auth, error) {
	stmt, err := u.db.Prepare(utils.LOGIN)
	user := models.Auth{}
	if err != nil {
		return &user, err
	}
	errQuery := stmt.QueryRow(username, password).Scan(user.UserID, user.UserLevelID, user.UserStatus)
	if errQuery != nil {
		return nil, err
	}
	defer stmt.Close()
	return &user, nil
}

//login 2
func (u *UserRepoImpl) ReadUserByUsername(username string) (*models.Auth, error) {
	fmt.Println("MASUK REPO", username)
	stmt, err := u.db.Prepare(utils.SELECT_AUTH_BY_USERNAME)
	user := models.Auth{}
	if err != nil {
		fmt.Println(err)
		return &user, err
	}
	errQuery := stmt.QueryRow(username).Scan(&user.AuthID, &user.Username, &user.Password, &user.UserID, &user.UserLevelID, &user.UserStatus)
	log.Println(errQuery)
	if errQuery != nil {
		return &user, err
	}
	fmt.Println(user.Password)
	defer stmt.Close()
	return &user, nil
}
