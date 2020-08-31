package controllers

import (
	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	"net/http"
)

type UsersHandler struct {
	UserUsecases usecases.UsersUseCases
}

func UsersControll(r *mux.Router, service usecases.UsersUseCases) {
	UsersHandler := UsersHandler{service}
	r.HandleFunc("/auth", UsersHandler.Auth).Methods(http.MethodPost)

}

func (u UsersHandler) Auth(w http.ResponseWriter, r *http.Request) {

}
