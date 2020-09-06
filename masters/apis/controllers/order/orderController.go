package orderControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	orderUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/order"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type OrderHandler struct {
	orderUsecase orderUsecases.OrderUsecase
}

func InitOrderController(orderUsecase orderUsecases.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

func (o *OrderHandler) OrderAPI(r *mux.Router) {
	orders := r.PathPrefix("/orders").Subrouter()
	orders.HandleFunc("", o.ListAllOrder).Methods(http.MethodGet)

	order := r.PathPrefix("/order").Subrouter()
	order.HandleFunc("/{id}", o.GetOrderByID).Methods(http.MethodGet)
	order.HandleFunc("/add", o.AddOrder).Methods(http.MethodPost)
	order.HandleFunc("/update/{id}", o.UpdateOrderPaid).Methods(http.MethodPut)
	order.HandleFunc("/delete/{id}", o.UpdateOrderCancel).Methods(http.MethodDelete)
}

func (o *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	storeID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&order, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = o.orderUsecase.AddOrder(storeID, &order)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			order, err := o.orderUsecase.GetOrderByID(order.OrderID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, order)
			}
		}
	}
}

func (o *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	orderID := utils.DecodePathVariabel("id", r)
	order, err := o.orderUsecase.GetOrderByID(orderID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, order)
	}
}

func (o *OrderHandler) ListAllOrder(w http.ResponseWriter, r *http.Request) {
	storeID := utils.DecodePathVariabel("id", r)
	orders, err := o.orderUsecase.GetAllOrderByStore(storeID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, orders)
	}
}

func (o *OrderHandler) UpdateOrderPaid(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	orderID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&order, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = o.orderUsecase.UpdateOrderPaid(orderID, &order)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			order, err := o.orderUsecase.GetOrderByID(orderID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, order)
			}
		}
	}
}

func (o *OrderHandler) UpdateOrderCancel(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	orderID := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&payment, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = o.orderUsecase.UpdateOrderCancel(orderID, &payment)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			order, err := o.orderUsecase.GetOrderByID(orderID)
			if err != nil {
				log.Print(err)
			} else {
				utils.HandleResponse(w, http.StatusOK, order)
			}
		}
	}
}
