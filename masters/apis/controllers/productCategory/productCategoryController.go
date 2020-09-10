package productCategoryControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	productCategoryUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/productCategory"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type ProductCategoryHandler struct {
	ProductCategoryUsecase productCategoryUsecases.ProductCategoryUsecase
}

func ProductCategoryController(productCategoryUsecase productCategoryUsecases.ProductCategoryUsecase) *ProductCategoryHandler {
	return &ProductCategoryHandler{ProductCategoryUsecase: productCategoryUsecase}
}

func (pc *ProductCategoryHandler) ProductCategoryAPI(r *mux.Router) {
	productCategories := r.PathPrefix("/productCategories").Subrouter()
	productCategories.HandleFunc("", pc.ListAllProductCategory).Methods(http.MethodGet)

	productCategory := r.PathPrefix("/productCategory").Subrouter()
	productCategory.HandleFunc("/{id}", pc.GetProductCategoryByID).Methods(http.MethodGet)
	productCategory.HandleFunc("/add", pc.AddProductCategory).Methods(http.MethodPost)
	productCategory.HandleFunc("/update/{id}", pc.UpdateProductCategory).Methods(http.MethodPut)
	productCategory.HandleFunc("/delete/{id}", pc.DeleteProductCategory).Methods(http.MethodDelete)
}

func (pc *ProductCategoryHandler) AddProductCategory(w http.ResponseWriter, r *http.Request) {
	var productCategory models.ProductCategory
	err := utils.JsonDecoder(&productCategory, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = pc.ProductCategoryUsecase.AddProductCategory(&productCategory)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			productCategory, err := pc.ProductCategoryUsecase.GetProductCategoryByID(productCategory.ProductCategoryID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, productCategory)
			}
		}
	}
}

func (pc *ProductCategoryHandler) GetProductCategoryByID(w http.ResponseWriter, r *http.Request) {
	productCategoryID := utils.DecodePathVariabel("id", r)
	productCategory, err := pc.ProductCategoryUsecase.GetProductCategoryByID(productCategoryID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, productCategory)
	}
}

func (pc *ProductCategoryHandler) ListAllProductCategory(w http.ResponseWriter, r *http.Request) {
	productCategories, err := pc.ProductCategoryUsecase.GetAllProductCategory()
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, productCategories)
	}
}

func (pc *ProductCategoryHandler) UpdateProductCategory(w http.ResponseWriter, r *http.Request) {
	var productCategory models.ProductCategory
	productCategoryID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&productCategory, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = pc.ProductCategoryUsecase.UpdateProductCategory(productCategoryID, &productCategory)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			productCategory, err := pc.ProductCategoryUsecase.GetProductCategoryByID(productCategoryID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, productCategory)
			}
		}
	}
}

func (pc *ProductCategoryHandler) DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) > 0 {
		err := pc.ProductCategoryUsecase.DeleteProductCategory(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
}
