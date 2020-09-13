package models

type Wallet struct {
	WalletID string `json:"walletID"`
	UserID   string `json:"userID"`
	Amount   string `json:"amount"`
}

type TopUp struct {
	TopUpID       string `json:"topUpID"`
	TopUpAmount   string `json:"topUpAmount"`
	UserID        string `json:"userID"`
	UserFirstName string `json:"userFirstName"`
	TopUpDate     string `json:"topUpDate"`
	TopUpStatus   string `json:"topUpStatus"`
}
