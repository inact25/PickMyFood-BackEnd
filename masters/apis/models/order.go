package models

type Order struct {
	OrderID      string      `json:"orderID"`
	OrderCreated string      `json:"orderCreated"`
	StoreID      string      `json:"storeID"`
	SoldItems    []SoldItems `json:"soldItems"`
}

type SoldItems struct {
	Qty               string `json;"qty"`
	ProductID         string `json:"productID"`
	ProductName       string `json:"productName"`
	UserID            string `json:"userID"`
	UserFirstName     string `json:"userFirstName"`
	Price             string `json:"price"`
	Subtotal          string `json:"subtotal"`
	OrderDetailStatus string `json:"orderDetailStatus"`
}
