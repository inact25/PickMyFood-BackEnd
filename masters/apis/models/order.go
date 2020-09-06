package models

type Order struct {
	OrderID      string      `json:"orderId"`
	OrderCreated string      `json:"orderCreated"`
	OrderDetail  OrderDetail `json:"orderDetail"`
	Store        Store       `json;"store"`
}

type OrderDetail struct {
	Qty               int        `json;"qty"`
	Product           []*Product `json:"product"`
	User              User       `json:"user"`
	Price             string     `json:"price"`
	OrderDetailStatus string     `json:"orderDetailStatus"`
}

type Payment struct {
	TransactionID      string `json:"transactionID"`
	Amount             string `json:"amount"`
	TransactionCreated string `json:"transactionCreated"`
	TransactionStatus  string `json:"transactionStatus"`
	Order              Order  `json:"order"`
}
