package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type UsersHandler struct {
	UserUsecases usecases.UserUseCase
}

func UsersController(UserUsecases usecases.UserUseCase) *UsersHandler {
	return &UsersHandler{UserUsecases: UserUsecases}
}

func (u *UsersHandler) Authenticate(r *mux.Router) {
	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("", u.ListAllUser).Methods(http.MethodGet)
	user.HandleFunc("/{id}", u.GetUserId).Methods(http.MethodGet)
	user.HandleFunc("/register", u.Register).Methods(http.MethodPost)
	user.HandleFunc("/login", u.Login).Methods(http.MethodPost)
	user.HandleFunc("/update/{id}", u.UpdateUser).Methods(http.MethodPut)
	user.HandleFunc("/delete/{id}", u.DeleteUser).Methods(http.MethodDelete)
}

func (u *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := utils.JsonDecoder(&user, r)
	user.Auth.Password = utils.Encrypt([]byte(user.Auth.Password))
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = u.UserUsecases.AddUser(&user)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			user, err := u.UserUsecases.GetUserByID(user.UserID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, user)
			}
		}
	}
}
func (u *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var auth models.Auth
	err := utils.JsonDecoder(&auth, r)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		authTemp, err := u.UserUsecases.ReadUserByUsername(auth.Username)
		if err != nil {
			utils.HandleResponseError(w, http.StatusBadGateway, utils.BAD_GATEWAY)
		}
		fmt.Println(auth.UserID)
		fmt.Println(auth.Username)
		fmt.Println(authTemp.Password)
		fmt.Println(auth.Password)

		isValid := utils.CompareEncrypt(authTemp.Password, []byte(auth.Password))
		fmt.Println(isValid)
		if isValid {
			token, err := utils.JwtEncoder(authTemp.Username, "Rahasia")
			// fmt.Println("Berhasil Login")
			// w.Write([]byte("Berhasil Login"))
			if err != nil {
				utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
			}
			authTemp.Token = models.Token{Key: token}
			utils.HandleResponse(w, http.StatusOK, authTemp)
		} else {
			utils.HandleResponseError(w, http.StatusUnauthorized, "Wrong password or username")
		}
	}
}

func (u *UsersHandler) GetUserId(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	user, err := u.UserUsecases.GetUserByID(id)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, user)
	}
}

func (u *UsersHandler) ListAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserUsecases.GetAllUser()
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, users)
	}
}

func (u *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&user, r)
	user.Auth.Password = utils.Encrypt([]byte(user.Auth.Password))
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = u.UserUsecases.UpdateUser(id, &user)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			user, err := u.UserUsecases.GetUserByID(id)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, user)
			}
		}
	}
}

func (u *UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) > 0 {
		err := u.UserUsecases.DeleteUser(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
}
