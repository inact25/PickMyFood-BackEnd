package models

type Store struct {
	StoreID       string        `json:"storeID"`
	StoreName     string        `json:"storeName"`
	StoreAddress  string        `json:"storeAddress"`
	StoreOwner    string        `json:"storeOwner"`
	StoreStatus   string        `json:"storeStatus"`
	StoreUsername string        `json:"storeUsername"`
	StorePassword string        `json;"storePassword"`
	StoreImage    string        `json:"storeImage"`
	StoreCategory StoreCategory `json:"storeCategory"`
	Token         Token         `json:"token"`
}

type StoreCategory struct {
	StoreCategoryID   string `json:"storeCategoryID"`
	StoreCategoryName string `json:"storeCategoryName"`
}
