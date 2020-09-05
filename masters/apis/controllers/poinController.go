package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	"github.com/inact25/PickMyFood-BackEnd/utils/message"
	"github.com/inact25/PickMyFood-BackEnd/utils/tools"
)

type PointsHandler struct {
	poinUsecases usecases.PoinUseCases
}

func PointsController(r *mux.Router, service usecases.PoinUseCases) {
	PointsHandler := PointsHandler{service}
	r.HandleFunc("/points", PointsHandler.GetPoints).Methods(http.MethodGet)
	r.HandleFunc("/point/{sid}", PointsHandler.GetPointByID).Methods(http.MethodGet)
	r.HandleFunc("/point", PointsHandler.PostPoint()).Methods(http.MethodPost)
	r.HandleFunc("/point/{sid}", PointsHandler.UpdatePoint()).Methods(http.MethodPut)
	r.HandleFunc("/point/{sid}", PointsHandler.DeletePoint()).Methods(http.MethodDelete)

}

func (s *PointsHandler) GetPoints(w http.ResponseWriter, r *http.Request) {
	points, err := s.poinUsecases.GetPoints()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfPoints, err := json.Marshal(points)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfPoints))
	w.Write([]byte("Data successfully found"))
}

func (s *PointsHandler) GetPointByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	points, err := s.poinUsecases.GetPointByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfPoints, err := json.Marshal(points)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfPoints))
}

func (s *PointsHandler) PostPoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.PoinModels
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.poinUsecases.PostPoint(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *PointsHandler) UpdatePoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.PoinModels
		tools.Parser(r, &data)

		result, err := s.poinUsecases.UpdatePoint(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *PointsHandler) DeletePoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.poinUsecases.DeletePoint(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

}
