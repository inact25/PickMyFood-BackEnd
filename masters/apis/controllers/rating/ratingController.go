package ratingControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	ratingUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/rating"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type RatingsHandler struct {
	ratingUsecases ratingUsecases.RatingUseCases
}

func RatingController(ratingUsecases ratingUsecases.RatingUseCases) *RatingsHandler {
	return &RatingsHandler{ratingUsecases: ratingUsecases}
}

func (s *RatingsHandler) RatingAPI(r *mux.Router) {
	//	RatingsHandler := RatingsHandler{service}
	r.HandleFunc("/ratings", s.GetRatings).Methods(http.MethodGet)
	r.HandleFunc("/rating/{sid}", s.GetRatingByID).Methods(http.MethodGet)
	r.HandleFunc("/rating/post", s.PostRating).Methods(http.MethodPost)
	r.HandleFunc("/rating/update/{sid}", s.UpdateRating).Methods(http.MethodPut)
	r.HandleFunc("/rating/delete/{sid}", s.DeleteRating).Methods(http.MethodDelete)

}

func (s *RatingsHandler) GetRatings(w http.ResponseWriter, r *http.Request) {
	ratings, err := s.ratingUsecases.GetRatings()
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, ratings)
	}
}

func (s *RatingsHandler) GetRatingByID(w http.ResponseWriter, r *http.Request) {
	ratingID := utils.DecodePathVariabel("sid", r)
	rating, err := s.ratingUsecases.GetRatingByID(ratingID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, rating)
	}
}

func (s *RatingsHandler) PostRating(w http.ResponseWriter, r *http.Request) {
	var rating models.RatingModels
	// id := utils.DecodePathVariabel("sid", r)
	err := utils.JsonDecoder(&rating, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.ratingUsecases.PostRating(&rating)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, rating)
		}
	}
}

func (s *RatingsHandler) UpdateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.RatingModels
	id := utils.DecodePathVariabel("sid", r)
	err := utils.JsonDecoder(&rating, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.ratingUsecases.UpdateRating(&rating, id)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, rating)
		}
	}
}

func (s *RatingsHandler) DeleteRating(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("sid", r)
	if len(id) > 0 {
		err := s.ratingUsecases.DeleteRating(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}

}
