package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
)

type StoresHandler struct {
	storeUsecases usecases.StoresUseCases
}

func StoresController(r *mux.Router, service usecases.StoresUseCases) {
	StoresHandler := StoresHandler{service}
	r.HandleFunc("/store", StoresHandler.GetStores).Methods(http.MethodGet)
	r.HandleFunc("/store/{sid}", StoresHandler.GetStoreByID).Methods(http.MethodGet)
	r.HandleFunc("/store/{sid}", StoresHandler.DeleteStore).Methods(http.MethodDelete)

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

func (s *StoresHandler) DeleteStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	stores, err := s.storeUsecases.DeleteStore(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfStores, err := json.Marshal(stores)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfStores))
	w.Write([]byte("Data successfully deleted"))
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
