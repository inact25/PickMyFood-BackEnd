package models

type Payment struct {
	TransactionID      string `json:"transactionID"`
	OrderID            string `json:"orderID"`
	UserID             string `json:"userID"`
	UserFirstName      string `json:"userFirstName"`
	Amount             string `json:"amount"`
	TransactionCreated string `json:"transactionCreated"`
	TransactionStatus  string `json:"transactionStatus"`
}
