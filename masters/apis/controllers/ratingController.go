package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
)

type RatingsHandler struct {
	ratingUsecases usecases.RatingUseCases
}

func RatingController(r *mux.Router, service usecases.RatingUseCases) {
	RatingsHandler := RatingsHandler{service}
	r.HandleFunc("/ratings", RatingsHandler.GetRatings).Methods(http.MethodGet)
	r.HandleFunc("/rating/{sid}", RatingsHandler.GetRatingByID).Methods(http.MethodGet)

	// r.HandleFunc("/product_price", ProductsHandler.GetProductsPrice).Methods(http.MethodGet)

	// r.HandleFunc("/product_category", ProductsHandler.GetProductsCategory).Methods(http.MethodGet)

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
