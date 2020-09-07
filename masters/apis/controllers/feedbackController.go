package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	"github.com/inact25/PickMyFood-BackEnd/utils"
	"github.com/inact25/PickMyFood-BackEnd/utils/message"
	"github.com/inact25/PickMyFood-BackEnd/utils/tools"
)

type FeedbacksHandler struct {
	feedbackUsecases usecases.FeedbackUseCases
}

func FeedbacksController(r *mux.Router, service usecases.FeedbackUseCases) {
	FeedbacksHandler := FeedbacksHandler{service}
	r.HandleFunc("/feedbacks", FeedbacksHandler.GetFeedbacks).Methods(http.MethodGet)
	r.HandleFunc("/feedback/{sid}", FeedbacksHandler.GetFeedbackByID).Methods(http.MethodGet)
	r.HandleFunc("/feedback", FeedbacksHandler.PostFeedback()).Methods(http.MethodPost)
	r.HandleFunc("/feedback/{sid}", FeedbacksHandler.UpdateFeedback()).Methods(http.MethodPut)
	r.HandleFunc("/feedback/{sid}", FeedbacksHandler.DeleteFeedback).Methods(http.MethodDelete)

}

func (s *FeedbacksHandler) GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := s.feedbackUsecases.GetFeedbacks()
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

func (s *FeedbacksHandler) PostFeedback() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.FeedbackModels
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.feedbackUsecases.PostFeedback(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *FeedbacksHandler) UpdateFeedback() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.FeedbackModels
		tools.Parser(r, &data)

		result, err := s.feedbackUsecases.UpdateFeedback(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
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
