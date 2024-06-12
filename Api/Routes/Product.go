package routes

import (
	controllers "crud/Api/Controllers"
	usecase "crud/Application/UseCase"
	domain "crud/Domain"

	"github.com/gorilla/mux"
)

func RegisterRoutesProduct(r *mux.Router, repositoryMysql domain.ProductRepository, repositoryRedis domain.ProductRepository) {
	productUseCase := usecase.NewProductUseCase(repositoryMysql, repositoryRedis)
	productControllers := controllers.NewProductControllers(productUseCase)
	r.HandleFunc("product/{id}", productControllers.GetProductByID).Methods("GET")
}
