package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type PointsHandler struct {
	poinUsecases usecases.PoinUseCases
}

func PointsController(poinUsecases usecases.PoinUseCases) *PointsHandler {
	return &PointsHandler{poinUsecases: poinUsecases}
}

func (s *PointsHandler) PointAPI(r *mux.Router) {
	//PointsHandler := PointsHandler{service}
	r.HandleFunc("/points", s.GetPoints).Methods(http.MethodGet)
	r.HandleFunc("/point/{sid}", s.GetPointByID).Methods(http.MethodGet)
	r.HandleFunc("/point/post", s.PostPoint).Methods(http.MethodPost)
	r.HandleFunc("/point/update/{sid}", s.UpdatePoint).Methods(http.MethodPut)
	r.HandleFunc("/point/delete/{sid}", s.DeletePoint).Methods(http.MethodDelete)

	r.HandleFunc("/point/update_user_point/{sid}", s.UpdateUserPoint).Methods(http.MethodPut)
}

func (s *PointsHandler) GetPoints(w http.ResponseWriter, r *http.Request) {
	points, err := s.poinUsecases.GetPoints()
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, points)
	}
}

func (s *PointsHandler) GetPointByID(w http.ResponseWriter, r *http.Request) {
	pointID := utils.DecodePathVariabel("sid", r)
	point, err := s.poinUsecases.GetPointByID(pointID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, point)
	}
}

func (s *PointsHandler) PostPoint(w http.ResponseWriter, r *http.Request) {
	var point models.PoinModels
	id := utils.DecodePathVariabel("sid", r)
	err := utils.JsonDecoder(&point, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.poinUsecases.PostPoint(&point, id)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, point)
		}
	}
}

func (s *PointsHandler) UpdatePoint(w http.ResponseWriter, r *http.Request) {
	var point models.PoinModels
	id := utils.DecodePathVariabel("sid", r)
	err := utils.JsonDecoder(&point, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.poinUsecases.UpdatePoint(&point, id)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, point)
		}
	}
}

func (s *PointsHandler) DeletePoint(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("sid", r)
	if len(id) > 0 {
		err := s.poinUsecases.DeletePoint(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}

}

func (s *PointsHandler) UpdateUserPoint(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := utils.DecodePathVariabel("sid", r)
	err := utils.JsonDecoder(&user, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.poinUsecases.UpdateUserPoint(id, &user)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, user)
		}
	}
}
