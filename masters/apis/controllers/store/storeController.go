package storeControllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	storeusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/store"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type StoreHandler struct {
	storeUsecase storeusecases.StoreUsecase
}

func StoreController(storeUsecase storeusecases.StoreUsecase) *StoreHandler {
	return &StoreHandler{storeUsecase: storeUsecase}
}

func (s *StoreHandler) StoreAPI(r *mux.Router) {
	stores := r.PathPrefix("/stores").Subrouter()
	stores.HandleFunc("", s.ListAllStore).Methods(http.MethodGet)

	store := r.PathPrefix("/store").Subrouter()
	store.HandleFunc("/{id}", s.GetStoreByID).Methods(http.MethodGet)
	store.HandleFunc("/register", s.RegisterStore).Methods(http.MethodPost)
	store.HandleFunc("/login", s.LoginStore).Methods(http.MethodPost)
	store.HandleFunc("/update/{id}", s.UpdateStore).Methods(http.MethodPut)
	store.HandleFunc("/delete/{id}", s.DeleteStore).Methods(http.MethodDelete)
}

func (s *StoreHandler) RegisterStore(w http.ResponseWriter, r *http.Request) {

	var store models.Store
	err := utils.JsonDecoder(&store, r)
	store.StorePassword = utils.Encrypt([]byte(store.StorePassword))
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.storeUsecase.AddStore(&store)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			store, err := s.storeUsecase.GetStoreByID(store.StoreID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, store)
			}
		}
	}
}
func (s *StoreHandler) LoginStore(w http.ResponseWriter, r *http.Request) {
	var store models.Store
	err := utils.JsonDecoder(&store, r)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		storeTemp, err := s.storeUsecase.Auth(store.StoreUsername)
		if err != nil {
			utils.HandleResponseError(w, http.StatusBadGateway, utils.BAD_GATEWAY)
		}
		fmt.Println(store.StoreID)
		fmt.Println(store.StoreUsername)
		fmt.Println(storeTemp.StorePassword)
		fmt.Println(store.StorePassword)

		isValid := utils.CompareEncrypt(storeTemp.StorePassword, []byte(store.StorePassword))
		fmt.Println(isValid)
		if isValid {
			token, err := utils.JwtEncoder(storeTemp.StoreUsername, "Rahasia")
			if err != nil {
				utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
			}
			storeTemp.Token = models.Token{Key: token}
			utils.HandleResponse(w, http.StatusOK, storeTemp)
		} else {
			utils.HandleResponseError(w, http.StatusUnauthorized, "Wrong password or username")
		}
	}
}

func (s *StoreHandler) GetStoreByID(w http.ResponseWriter, r *http.Request) {
	storeID := utils.DecodePathVariabel("id", r)
	store, err := s.storeUsecase.GetStoreByID(storeID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, store)
	}
}

func (s *StoreHandler) ListAllStore(w http.ResponseWriter, r *http.Request) {
	stores, err := s.storeUsecase.GetAllStore()
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, stores)
	}
}

func (s *StoreHandler) UpdateStore(w http.ResponseWriter, r *http.Request) {
	var store models.Store
	storeID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&store, r)
	store.StorePassword = utils.Encrypt([]byte(store.StorePassword))
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = s.storeUsecase.UpdateStore(storeID, &store)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			store, err := s.storeUsecase.GetStoreByID(storeID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, store)
			}
		}
	}
}

func (s *StoreHandler) DeleteStore(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) > 0 {
		err := s.storeUsecase.DeleteStore(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
}
