package models

type Wallet struct {
	WalletID string `json:"walletID"`
	UserID   string `json:"userID"`
	Amount   string `json:"amount"`
	User     User   `json:"user"`
}

type TopUp struct {
	TopUpID     string `json:"topUpID"`
	TopUpAmount string `json:"topUpAmount"`
	userID      string `json:"userID"`
	TopUpDate   string `json:"topUpDate"`
	Status      string `json:"status"`
}
