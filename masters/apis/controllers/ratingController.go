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

type RatingsHandler struct {
	ratingUsecases usecases.RatingUseCases
}

func RatingController(ratingUsecases usecases.RatingUseCases) *RatingsHandler {
	return &RatingsHandler{ratingUsecases: ratingUsecases}
}

func (s *RatingsHandler) RatingAPI(r *mux.Router) {
	//	RatingsHandler := RatingsHandler{service}
	r.HandleFunc("/ratings", s.GetRatings).Methods(http.MethodGet)
	r.HandleFunc("/rating/{sid}", s.GetRatingByID).Methods(http.MethodGet)
	r.HandleFunc("/rating", s.PostRating()).Methods(http.MethodPost)
	r.HandleFunc("/rating/{sid}", s.UpdateRating()).Methods(http.MethodPut)
	r.HandleFunc("/rating/{sid}", s.DeleteRating()).Methods(http.MethodDelete)

}

func (s *RatingsHandler) GetRatings(w http.ResponseWriter, r *http.Request) {
	ratings, err := s.ratingUsecases.GetRatings()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfRatings, err := json.Marshal(ratings)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfRatings))
	w.Write([]byte("Data successfully found"))
}

func (s *RatingsHandler) GetRatingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	ratings, err := s.ratingUsecases.GetRatingByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfRatings, err := json.Marshal(ratings)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfRatings))
}

func (s *RatingsHandler) PostRating() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.RatingModels
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.ratingUsecases.PostRating(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *RatingsHandler) UpdateRating() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.RatingModels
		tools.Parser(r, &data)

		result, err := s.ratingUsecases.UpdateRating(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *RatingsHandler) DeleteRating() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.ratingUsecases.DeleteRating(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

}
