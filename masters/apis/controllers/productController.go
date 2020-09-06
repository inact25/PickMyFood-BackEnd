package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis/usecases"
	"github.com/inact25/PickMyFood-BackEnd/utils/message"
	"github.com/inact25/PickMyFood-BackEnd/utils/tools"
)

type ProductsHandler struct {
	productUsecases usecases.ProductsUseCases
}

func ProductsController(r *mux.Router, service usecases.ProductsUseCases) {
	ProductsHandler := ProductsHandler{service}
	r.HandleFunc("/products", ProductsHandler.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{sid}", ProductsHandler.GetProductByID).Methods(http.MethodGet)
	r.HandleFunc("/product", ProductsHandler.PostProduct()).Methods(http.MethodPost)
	r.HandleFunc("/product/{sid}", ProductsHandler.UpdateProduct()).Methods(http.MethodPut)
	r.HandleFunc("/product/{sid}", ProductsHandler.DeleteProduct()).Methods(http.MethodDelete)

	r.HandleFunc("/products_price", ProductsHandler.GetProductsPrice).Methods(http.MethodGet)
	r.HandleFunc("/product_price/{sid}", ProductsHandler.GetProductPriceByID).Methods(http.MethodGet)
	r.HandleFunc("/product_price", ProductsHandler.PostProductPrice()).Methods(http.MethodPost)
	r.HandleFunc("/product_price/{sid}", ProductsHandler.UpdateProductPrice()).Methods(http.MethodPut)
	r.HandleFunc("/product_price/{sid}", ProductsHandler.DeleteProductPrice()).Methods(http.MethodDelete)

	r.HandleFunc("/products_category", ProductsHandler.GetProductsCategory).Methods(http.MethodGet)
	r.HandleFunc("/product_category/{sid}", ProductsHandler.GetProductCategoryByID).Methods(http.MethodGet)
	r.HandleFunc("/product_category", ProductsHandler.PostProductCategory()).Methods(http.MethodPost)
	r.HandleFunc("/product_category/{sid}", ProductsHandler.UpdateProductCategory()).Methods(http.MethodPut)
	r.HandleFunc("/product_category/{sid}", ProductsHandler.DeleteProductCategory()).Methods(http.MethodDelete)
}

func (s *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.productUsecases.GetProducts()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
	w.Write([]byte("Data successfully found"))
}

func (s *ProductsHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	products, err := s.productUsecases.GetProductByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
}

func (s *ProductsHandler) PostProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.ProductModels
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.productUsecases.PostProduct(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) UpdateProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.ProductModels
		tools.Parser(r, &data)

		result, err := s.productUsecases.UpdateProduct(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) DeleteProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.productUsecases.DeleteProduct(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

}

func (s *ProductsHandler) GetProductsPrice(w http.ResponseWriter, r *http.Request) {
	products, err := s.productUsecases.GetProductsPrice()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
	w.Write([]byte("Data successfully found"))
}

func (s *ProductsHandler) GetProductPriceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	productsPrice, err := s.productUsecases.GetProductPriceByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfProductsPrice, err := json.Marshal(productsPrice)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProductsPrice))
}

func (s *ProductsHandler) PostProductPrice() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.ProductPrice
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.productUsecases.PostProductPrice(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) UpdateProductPrice() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.ProductPrice
		tools.Parser(r, &data)

		result, err := s.productUsecases.UpdateProductPrice(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) DeleteProductPrice() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.productUsecases.DeleteProductPrice(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

}

func (s *ProductsHandler) GetProductsCategory(w http.ResponseWriter, r *http.Request) {
	products, err := s.productUsecases.GetProductsCategory()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfProducts, err := json.Marshal(products)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProducts))
	w.Write([]byte("Data successfully found"))
}

func (s *ProductsHandler) GetProductCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["sid"]

	productsCategory, err := s.productUsecases.GetProductCategoryByID(strID)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfProductsCategory, err := json.Marshal(productsCategory)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfProductsCategory))
}

func (s *ProductsHandler) PostProductCategory() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data models.ProductCategory
		tools.Parser(r, &data)

		fmt.Println(data)

		result, err := s.productUsecases.PostProductCategory(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Posting Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Posting Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) UpdateProductCategory() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data models.ProductCategory
		tools.Parser(r, &data)

		result, err := s.productUsecases.UpdateProductCategory(tools.GetPathVar("sid", r), data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Update Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Update Success", http.StatusOK, result))
	}
}

func (s *ProductsHandler) DeleteProductCategory() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := s.productUsecases.DeleteProductCategory(tools.GetPathVar("sid", r))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Response("Delete By ID Failed", http.StatusBadRequest, err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Response("Delete By ID Success", http.StatusOK, result))
	}

}
