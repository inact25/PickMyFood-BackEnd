package models

type OrderModels struct {
	OrderID      string `json:"order_id"`
	OrderCreated string `json:"order_created"`
	StoreID      string `json:"store_id"`
}

type OrderDetailModels struct {
	Qty               string `json:"qty"`
	OrderID           string `json:"order_id"`
	ProductID         string `json:"product_id"`
	UserID            string `json:"user_id"`
	Price             string `json:"price"`
	OrderDetailStatus string `json:"order_detail_status"`
}
