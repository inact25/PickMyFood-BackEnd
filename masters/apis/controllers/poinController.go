package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
)

type PointsHandler struct {
	poinUsecases usecases.PoinUseCases
}

func PointsController(r *mux.Router, service usecases.PoinUseCases) {
	PointsHandler := PointsHandler{service}
	r.HandleFunc("/points", PointsHandler.GetPoints).Methods(http.MethodGet)
	// r.HandleFunc("/product/{sid}", ProductsHandler.GetProductByID).Methods(http.MethodGet)

	// r.HandleFunc("/product_price", ProductsHandler.GetProductsPrice).Methods(http.MethodGet)

	// r.HandleFunc("/product_category", ProductsHandler.GetProductsCategory).Methods(http.MethodGet)

}

func (s *PointsHandler) GetPoints(w http.ResponseWriter, r *http.Request) {
	points, err := s.poinUsecases.GetPoints()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfPoints, err := json.Marshal(points)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfPoints))
	w.Write([]byte("Data successfully found"))
}

// func (s *ProductsHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	strID := vars["sid"]

// 	products, err := s.productUsecases.GetProductByID(strID)
// 	if err != nil {
// 		w.Write([]byte("Data Not Found!"))
// 	}

// 	byteOfProducts, err := json.Marshal(products)

// 	if err != nil {
// 		w.Write([]byte("Data not found!"))
// 	}

// 	w.Write([]byte(byteOfProducts))
// }
