package productControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	productUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/product"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type ProductHandler struct {
	productUsecase productUsecases.ProductUsecase
}

func InitProductController(productUsecase productUsecases.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (p *ProductHandler) ProductAPI(r *mux.Router) {
	products := r.PathPrefix("/products").Subrouter()
	products.HandleFunc("", p.ListAllProduct).Methods(http.MethodGet)

	product := r.PathPrefix("/product").Subrouter()
	product.HandleFunc("/{id}", p.GetProductByID).Methods(http.MethodGet)
	product.HandleFunc("/add/{id}", p.AddProduct).Methods(http.MethodPost)
	product.HandleFunc("/update/{id}", p.UpdateProduct).Methods(http.MethodPut)
	product.HandleFunc("/delete/{id}", p.DeleteProduct).Methods(http.MethodDelete)
}

func (p *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	storeID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&product, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = p.productUsecase.AddProduct(storeID, &product)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			product, err := p.productUsecase.GetProductByID(product.ProductID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, product)
			}
		}
	}
}

func (p *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := utils.DecodePathVariabel("id", r)
	product, err := p.productUsecase.GetProductByID(productID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, product)
	}
}

func (p *ProductHandler) ListAllProduct(w http.ResponseWriter, r *http.Request) {
	storeID := utils.DecodePathVariabel("id", r)
	products, err := p.productUsecase.GetAllProductByStore(storeID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, products)
	}
}

func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	productID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&product, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = p.productUsecase.UpdateProductWithPrice(productID, &product)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			product, err := p.productUsecase.GetProductByID(productID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, product)
			}
		}
	}
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	if len(id) > 0 {
		err := p.productUsecase.DeleteProduct(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
}
