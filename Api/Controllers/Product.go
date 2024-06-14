package controllers

import (
	middleware "crud/Api/Middleware"
	domain "crud/Domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductControllers struct {
	product domain.ProductInteractor
}

func NewProductControllers(productUseCase domain.ProductInteractor) *ProductControllers {
	return &ProductControllers{product: productUseCase}
}

func (productControllers *ProductControllers) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)
	GetID, err := strconv.Atoi(ID["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := productControllers.product.GetByID(GetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (ProductControllers *ProductControllers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.Create_By = r.Context().Value(middleware.UserIdKey).(int)
	err = ProductControllers.product.UpdateByID(product.ID, &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (ProductControllers *ProductControllers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)
	IntID, err := strconv.Atoi(ID["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ProductControllers.product.DeleteByID(IntID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
