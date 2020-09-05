package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
)

type ProductsHandler struct {
	productUsecases usecases.ProductsUseCases
}

func ProductsController(r *mux.Router, service usecases.ProductsUseCases) {
	ProductsHandler := ProductsHandler{service}
	r.HandleFunc("/products", ProductsHandler.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{sid}", ProductsHandler.GetProductByID).Methods(http.MethodGet)

	r.HandleFunc("/product_price", ProductsHandler.GetProductsPrice).Methods(http.MethodGet)

	r.HandleFunc("/product_category", ProductsHandler.GetProductsCategory).Methods(http.MethodGet)

}

func (s *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.productUsecases.GetProducts()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
	w.Write([]byte("Data successfully found"))
}

func (s *ProductsHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	products, err := s.productUsecases.GetProductByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
}

func (s *ProductsHandler) GetProductsPrice(w http.ResponseWriter, r *http.Request) {
	products, err := s.productUsecases.GetProductsPrice()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
	w.Write([]byte("Data successfully found"))
}

func (s *ProductsHandler) GetProductsCategory(w http.ResponseWriter, r *http.Request) {
	products, err := s.productUsecases.GetProductsCategory()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
	w.Write([]byte("Data successfully found"))
}
