package models

type StoreModels struct {
	StoreID         string `json:"store_id"`
	StoreName       string `json:"store_name"`
	StoreCategoryID string `json:"store_category_id"`
	StoreAddress    string `json:"store_address"`
	StoreOwner      string `json:"store_owner"`
	StoreStatus     string `json:"store_status"`
	StorePassword   string `json:"store_password"`
	StoreImages     string `json:"store_images"`
}
