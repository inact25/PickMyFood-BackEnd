package models

type Transaction struct {
	TransactionID      string `json:"transactionID"`
	Amount             string `json:"amount"`
	TransactionCreated string `json:"transactionCreated"`
	TransactionStatus  string `json:"transactionStatus"`
	Order              Order  `json:"order"`
}
