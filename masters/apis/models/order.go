package models

type Order struct {
	OrderID      string      `json:"orderId"`
	OrderCreated string      `json:"orderCreated"`
	StoreID      string      `json:"storeID"`
	SoldItems    []SoldItems `json:"soldItems"`
}

type SoldItems struct {
	Qty               string `json;"qty"`
	ProductID         string `json:"productID"`
	UserID            string `json:"userId"`
	Price             string `json:"price"`
	OrderDetailStatus string `json:"orderDetailStatus"`
}

type Payment struct {
	TransactionID      string `json:"transactionID"`
	Amount             string `json:"amount"`
	TransactionCreated string `json:"transactionCreated"`
	TransactionStatus  string `json:"transactionStatus"`
	Order              Order  `json:"order"`
}
