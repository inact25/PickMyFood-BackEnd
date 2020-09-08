package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type FeedbacksHandler struct {
	feedbackUsecases usecases.FeedbackUseCases
}

func FeedbacksController(r *mux.Router, service usecases.FeedbackUseCases) {
	FeedbacksHandler := FeedbacksHandler{service}
	r.HandleFunc("/feedbacks", FeedbacksHandler.GetFeedbacks).Methods(http.MethodGet)
	r.HandleFunc("/feedback/{sid}", FeedbacksHandler.GetFeedbackByID).Methods(http.MethodGet)
	r.HandleFunc("/feedback/post", FeedbacksHandler.PostFeedback).Methods(http.MethodPost)
	r.HandleFunc("/feedback/update/{sid}", FeedbacksHandler.UpdateFeedback).Methods(http.MethodPut)
	r.HandleFunc("/feedback/delete/{sid}", FeedbacksHandler.DeleteFeedback).Methods(http.MethodDelete)

}

func (s *FeedbacksHandler) GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := s.feedbackUsecases.GetFeedbacks()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, feedbacks)
	}
}

func (s *FeedbacksHandler) GetFeedbackByID(w http.ResponseWriter, r *http.Request) {
	feedbackID := utils.DecodePathVariabel("sid", r)
	feedback, err := s.feedbackUsecases.GetFeedbackByID(feedbackID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, feedback)
	}
}

func (s *FeedbacksHandler) PostFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.FeedbackModels
	id := utils.DecodePathVariabel("sid", r)
	err := utils.JsonDecoder(&feedback, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.feedbackUsecases.PostFeedback(&feedback, id)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, feedback)
		}
	}
}

func (s *FeedbacksHandler) UpdateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.FeedbackModels
	id := utils.DecodePathVariabel("sid", r)
	err := utils.JsonDecoder(&feedback, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.feedbackUsecases.UpdateFeedback(&feedback, id)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, feedback)
		}
	}

}

func (s *FeedbacksHandler) DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("sid", r)
	if len(id) > 0 {
		err := s.feedbackUsecases.DeleteFeedback(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
}
