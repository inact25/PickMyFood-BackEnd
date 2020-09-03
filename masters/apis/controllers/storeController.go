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

type StoresHandler struct {
	storeUsecases usecases.StoresUseCases
}

func StoresController(r *mux.Router, service usecases.StoresUseCases) {
	StoresHandler := StoresHandler{service}
	r.HandleFunc("/store", StoresHandler.GetStores).Methods(http.MethodGet)
	r.HandleFunc("/store/{sid}", StoresHandler.GetStoreByID).Methods(http.MethodGet)
	r.HandleFunc("/store", StoresHandler.PostStore()).Methods(http.MethodPost)
	r.HandleFunc("/store/{sid}", StoresHandler.UpdateStore()).Methods(http.MethodPut)
	r.HandleFunc("/store/{sid}", StoresHandler.DeleteStore()).Methods(http.MethodDelete)

	r.HandleFunc("/store_category", StoresHandler.GetStoreCategory).Methods(http.MethodGet)

}

func (s *StoresHandler) GetStores(w http.ResponseWriter, r *http.Request) {
	stores, err := s.storeUsecases.GetStores()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfStores, err := json.Marshal(stores)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfStores))
	w.Write([]byte("Data successfully found"))
}

func (s *StoresHandler) GetStoreByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	stores, err := s.storeUsecases.GetStoreByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfStores, err := json.Marshal(stores)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfStores))
}

func (s *StoresHandler) PostStore() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.StoreModels
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.storeUsecases.PostStore(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *StoresHandler) UpdateStore() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.StoreModels
		tools.Parser(r, &data)

		result, err := s.storeUsecases.UpdateStore(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *StoresHandler) DeleteStore() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.storeUsecases.DeleteStore(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

}

func (s *StoresHandler) GetStoreCategory(w http.ResponseWriter, r *http.Request) {
	stores, err := s.storeUsecases.GetStoresCategory()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfStores, err := json.Marshal(stores)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfStores))
	w.Write([]byte("Data successfully found"))
}
