package walletControllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/middlewares"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	walletusecases "github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases/wallet"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

type WalletHandler struct {
	WalletUsecases walletusecases.WalletUseCases
}

func WalletController(WalletUsecases walletusecases.WalletUseCases) *WalletHandler {
	return &WalletHandler{WalletUsecases: WalletUsecases}
}

func (wa *WalletHandler) WalletApi(r *mux.Router) {

	wallet := r.PathPrefix("/wallet").Subrouter()
	wallet.HandleFunc("/{id}", wa.WalletByID).Methods(http.MethodGet)
	wallet.HandleFunc("/topUp/{id}", wa.WalletTopUp).Methods(http.MethodPost)
	wallet.HandleFunc("/transfer/{id}", wa.WalletUpdateAmount).Methods(http.MethodPost)

	topUp := r.PathPrefix("/topUp").Subrouter()
	topUp.Use(middlewares.TokenValidationMiddleware)
	topUp.HandleFunc("", wa.ListTopUp).Methods(http.MethodGet)

	topUpuser := r.PathPrefix("/topUp").Subrouter()
	topUpuser.HandleFunc("/user/{id}", wa.GetTopUpByUser).Methods(http.MethodGet)
}

func (wa *WalletHandler) WalletByID(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	wallet, err := wa.WalletUsecases.GetWalletByID(id)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, wallet)
	}
}

//WalletTopUp
func (wa *WalletHandler) WalletTopUp(w http.ResponseWriter, r *http.Request) {
	var topUp models.TopUp
	id := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&topUp, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = wa.WalletUsecases.TopUpWallet(&topUp, id)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, topUp)
		}
	}
}

//UpdateAmount
func (wa *WalletHandler) WalletUpdateAmount(w http.ResponseWriter, r *http.Request) {
	var wallet models.Wallet
	id := utils.DecodePathVariabel("id", r)
	err := utils.JsonDecoder(&wallet, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = wa.WalletUsecases.UpdateAmountWallet(&wallet, id)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, wallet)
		}
	}
}

// all top up user
func (wa *WalletHandler) GetTopUpByUser(w http.ResponseWriter, r *http.Request) {
	id := utils.DecodePathVariabel("id", r)
	listTopUp, err := wa.WalletUsecases.GetAllTopUpByUser(id)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, listTopUp)
	}
}

// all top up
func (wa *WalletHandler) ListTopUp(w http.ResponseWriter, r *http.Request) {
	listTopUp, err := wa.WalletUsecases.GetAllTopUp()
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, listTopUp)
	}
}
