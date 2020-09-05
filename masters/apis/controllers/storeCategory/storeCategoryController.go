package storeCategoryControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	storeCategoryUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/storeCatergory"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type StoreCategoryHandler struct {
	storeCategoryUsecase storeCategoryUsecases.StoreCategoryUsecase
}

func StoreCategoryController(storeCategoryUsecase storeCategoryUsecases.StoreCategoryUsecase) *StoreCategoryHandler {
	return &StoreCategoryHandler{storeCategoryUsecase: storeCategoryUsecase}
}

func (sc *StoreCategoryHandler) StoreCategoryAPI(r *mux.Router) {
	storeCategories := r.PathPrefix("/storeCategories").Subrouter()
	storeCategories.HandleFunc("", sc.ListAllStoreCategory).Methods(http.MethodGet)

	storeCategory := r.PathPrefix("/storeCategory").Subrouter()
	storeCategory.HandleFunc("/{id}", sc.GetStoreCategoryByID).Methods(http.MethodGet)
	storeCategory.HandleFunc("/add", sc.AddStoreCategory).Methods(http.MethodPost)
	storeCategory.HandleFunc("/update/{id}", sc.UpdateStoreCategory).Methods(http.MethodPut)
	storeCategory.HandleFunc("/delete/{id}", sc.DeleteStoreCategory).Methods(http.MethodDelete)
}

func (sc *StoreCategoryHandler) AddStoreCategory(w http.ResponseWriter, r *http.Request) {
	var storeCategory models.StoreCategory
	err := utils.JsonDecoder(&storeCategory, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = sc.storeCategoryUsecase.AddStoreCategory(&storeCategory)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			storeCategory, err := sc.storeCategoryUsecase.GetStoreCategoryByID(storeCategory.StoreCategoryID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, storeCategory)
			}
		}
	}
}

func (sc *StoreCategoryHandler) GetStoreCategoryByID(w http.ResponseWriter, r *http.Request) {
	storeCategoryID := utils.DecodePathVariabel("id", r)
	storeCategory, err := sc.storeCategoryUsecase.GetStoreCategoryByID(storeCategoryID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, storeCategory)
	}
}

func (sc *StoreCategoryHandler) ListAllStoreCategory(w http.ResponseWriter, r *http.Request) {
	storeCategories, err := sc.storeCategoryUsecase.GetAllStoreCategory()
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, storeCategories)
	}
}

func (sc *StoreCategoryHandler) UpdateStoreCategory(w http.ResponseWriter, r *http.Request) {
	var storeCategory models.StoreCategory
	storeCategoryID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&storeCategory, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = sc.storeCategoryUsecase.UpdateStoreCategory(storeCategoryID, &storeCategory)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			storeCategory, err := sc.storeCategoryUsecase.GetStoreCategoryByID(storeCategoryID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, storeCategory)
			}
		}
	}
}

func (sc *StoreCategoryHandler) DeleteStoreCategory(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) > 0 {
		err := sc.storeCategoryUsecase.DeleteStoreCategory(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
}
