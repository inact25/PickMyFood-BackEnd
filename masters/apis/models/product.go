package models

type Product struct {
	ProductID       string          `json:"productID"`
	ProductName     string          `json:"productName"`
	ProductStock    string          `json:"productStock"`
	ProductStatus   string          `json:"productStatus"`
	ProductImage    string          `json:"productImage"`
	ProductPrice    ProductPrice    `json:"productPrice"`
	ProductCategory ProductCategory `json:"productCategory"`
	Store           Store           `json:"store"`
}

type ProductCategory struct {
	ProductCategoryID   string `json:"productCategoryID"`
	ProductCategoryName string `json:"productCategoryName"`
}

type ProductPrice struct {
	ProductPriceID string `json:"productPriceID"`
	Price          string `json:"price"`
	DateModified   string `json:"dateModified"`
}
