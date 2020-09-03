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

type ProductsHandler struct {
	productUsecases usecases.ProductsUseCases
}

func ProductsController(r *mux.Router, service usecases.ProductsUseCases) {
	ProductsHandler := ProductsHandler{service}
	r.HandleFunc("/products", ProductsHandler.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{sid}", ProductsHandler.GetProductByID).Methods(http.MethodGet)
	r.HandleFunc("/product", ProductsHandler.PostProduct()).Methods(http.MethodPost)
	r.HandleFunc("/product/{sid}", ProductsHandler.UpdateProduct()).Methods(http.MethodPut)
	r.HandleFunc("/product/{sid}", ProductsHandler.DeleteProduct()).Methods(http.MethodDelete)

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

func (s *ProductsHandler) PostProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.ProductModels
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.productUsecases.PostProduct(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) UpdateProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.ProductModels
		tools.Parser(r, &data)

		result, err := s.productUsecases.UpdateProduct(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) DeleteProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.productUsecases.DeleteProduct(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

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
