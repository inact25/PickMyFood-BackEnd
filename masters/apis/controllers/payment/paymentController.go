package paymentControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	paymentUsecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/payment"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type PaymentHandler struct {
	paymentUsecase paymentUsecases.PaymentUsecase
}

func InitPaymentController(paymentUsecase paymentUsecases.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUsecase: paymentUsecase}
}

func (p *PaymentHandler) PaymentAPI(r *mux.Router) {
	// payments := r.PathPrefix("/payments").Subrouter()
	// payments.HandleFunc("", p.ListAllPayment).Methods(http.MethodGet)

	payment := r.PathPrefix("/payment").Subrouter()
	// payment.HandleFunc("/{id}", p.GetPaymentByID).Methods(http.MethodGet)
	payment.HandleFunc("", p.PaymentWallet).Methods(http.MethodPost)
	transaction := r.PathPrefix("/transaction").Subrouter()
	transaction.HandleFunc("", p.UpdateTransaction).Queries("storeID", "{storeID}", "amount", "{amount}", "orderID", "{orderID}", "userID", "{userID}").Methods(http.MethodPost)
}

func (p *PaymentHandler) PaymentWallet(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	err := utils.JsonDecoder(&payment, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = p.paymentUsecase.PaymentWallet(&payment)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			// payment, err := o.paymentUsecase.GetpaymentByID(payment.paymentID)
			// if err != nil {
			// 	log.Print(err)
			// } else {
			utils.HandleResponse(w, http.StatusOK, payment)
			// }
		}
	}
}

func (p *PaymentHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var storeID string = mux.Vars(r)["storeID"]
	var amount string = mux.Vars(r)["amount"]
	var orderID string = mux.Vars(r)["orderID"]
	var userID string = mux.Vars(r)["userID"]

	err := p.paymentUsecase.UpdateTransaction(storeID, amount, orderID, userID)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, "Transaction Update Succesfully")
	}
}
