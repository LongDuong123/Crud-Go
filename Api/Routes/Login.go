package routes

import (
	controllers "crud/Api/Controllers"
	usecase "crud/Application/UseCase"
	domain "crud/Domain"

	"github.com/gorilla/mux"
)

func RegisterRoutesLogin(r *mux.Router, ur domain.UserRepository) {
	UseCaseLogin := usecase.NewLoginUseCase(ur)
	LoginController := controllers.NewLoginController(UseCaseLogin)
	r.HandleFunc("/login", LoginController.LoginUser).Methods("POST")
}
