package controllers

import (
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}
