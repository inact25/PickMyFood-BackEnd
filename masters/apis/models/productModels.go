package models

type ProductModels struct {
	ProductID         string `json:"product_id"`
	StoreID           string `json:"store_id"`
	ProductName       string `json:"product_name"`
	ProductCategoryID string `json:"product_category_id"`
	ProductStock      string `json:"product_stock"`
	ProductStatus     string `json:"product_status"`
}

type ProductCategory struct {
	ProductCategoryID   string `json:"product_category_id"`
	ProductCategoryName string `json:"product_category_name"`
}

type ProductPrice struct {
	ProductPriceID string `json:"product_category_id"`
	ProductID      string `json:"product_id"`
	Price          string `json:"price"`
	DateModified   string `json:"date_modified"`
}
