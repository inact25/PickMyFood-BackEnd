package userControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	userUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/user"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type UsersHandler struct {
	UserUsecases userUsecases.UserUseCase
}

func UsersController(UserUsecases userUsecases.UserUseCase) *UsersHandler {
	return &UsersHandler{UserUsecases: UserUsecases}
}

func (u *UsersHandler) Authenticate(r *mux.Router) {
	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/{id}", u.GetUserId).Methods(http.MethodGet)
	user.HandleFunc("/register", u.Register).Methods(http.MethodPost)
	user.HandleFunc("/login", u.Login).Methods(http.MethodPost)
	user.HandleFunc("/update/{id}", u.UpdateUser).Methods(http.MethodPut)
	user.HandleFunc("/delete/{id}", u.DeleteUser).Methods(http.MethodDelete)

	users := r.PathPrefix("/user").Subrouter()
	users.Use(middlewares.TokenValidationMiddleware)
	users.HandleFunc("", u.ListAllUser).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}").Methods(http.MethodGet)
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
		isValid := utils.CompareEncrypt(authTemp.Auth.Password, []byte(auth.Password))
		if isValid {
			token, err := utils.JwtEncoder(authTemp.Auth.Username, "Rahasia")
			if err != nil {
				utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
			}
			authTemp.Auth.Token = models.Token{Key: token}
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
	var page string = mux.Vars(r)["page"]
	var limit string = mux.Vars(r)["limit"]
	var keyword string = mux.Vars(r)["keyword"]

	users, err := u.UserUsecases.GetAllUser(keyword, page, limit)
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

//a
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
