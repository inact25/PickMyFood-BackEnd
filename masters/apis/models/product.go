package models

type Product struct {
	ProductID       string          `json:"productID"`
	ProductName     string          `json;"productName"`
	ProductStock    string          `json:"productStock"`
	ProductStatus   string          `json:"productStatus"`
	ProductCategory ProductCategory `json:"productCategory"`
	Store           Store           `json:"store"`
}

type ProductCategory struct {
	ProductCategoryID   string `json:"productCategoryID"`
	ProductCategoryName string `json:"productCategoryName"`
}
