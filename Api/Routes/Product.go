package routes

import (
	controllers "crud/Api/Controllers"
	middleware "crud/Api/Middleware"
	usecase "crud/Application/UseCase"
	domain "crud/Domain"

	"github.com/gorilla/mux"
)

func RegisterRoutesProduct(r *mux.Router, repositoryMysql domain.ProductRepository, repositoryRedis domain.ProductRepository) {
	ProductUseCase := usecase.NewProductUseCase(repositoryMysql, repositoryRedis)
	productControllers := controllers.NewProductControllers(ProductUseCase)
	r.HandleFunc("/product/{id}", productControllers.GetProductByID).Methods("GET")
	r.HandleFunc("/product", middleware.MiddleWare(productControllers.UpdateProduct)).Methods("PATCH")
	r.HandleFunc("/product/{id}", productControllers.DeleteProduct).Methods("DELETE")
}
