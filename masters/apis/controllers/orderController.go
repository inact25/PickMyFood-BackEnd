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

type OrdersHandler struct {
	orderUsecases usecases.OrderUseCases
}

func OrdersController(r *mux.Router, service usecases.OrderUseCases) {
	OrdersHandler := OrdersHandler{service}
	r.HandleFunc("/orders", OrdersHandler.GetOrders).Methods(http.MethodGet)
	r.HandleFunc("/order/{sid}", OrdersHandler.GetOrderByID).Methods(http.MethodGet)
	r.HandleFunc("/order", OrdersHandler.PostOrder()).Methods(http.MethodPost)
	r.HandleFunc("/order/{sid}", OrdersHandler.UpdateOrder()).Methods(http.MethodPut)
	r.HandleFunc("/order/{sid}", OrdersHandler.DeleteOrder()).Methods(http.MethodDelete)

}

func (s *OrdersHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := s.orderUsecases.GetOrders()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfOrders, err := json.Marshal(orders)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfOrders))
	w.Write([]byte("Data successfully found"))
}

func (s *OrdersHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	orders, err := s.orderUsecases.GetOrderByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfOrders, err := json.Marshal(orders)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfOrders))
}

func (s *OrdersHandler) PostOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.OrderModels
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.orderUsecases.PostOrder(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *OrdersHandler) UpdateOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.OrderModels
		tools.Parser(r, &data)

		result, err := s.orderUsecases.UpdateOrder(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *OrdersHandler) DeleteOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.orderUsecases.DeleteOrder(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

}
