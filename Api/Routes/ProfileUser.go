package routes

import (
	controllers "crud/Api/Controllers"
	middleware "crud/Api/Middleware"
	usecase "crud/Application/UseCase"
	domain "crud/Domain"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutesProfile(r *mux.Router, ur domain.UserRepository) {
	profileUseCase := usecase.NewProfileUseCase(ur)
	profileControllers := controllers.NewProfileControllers(profileUseCase)
	r.HandleFunc("/profile", middleware.MiddleWare(http.HandlerFunc(profileControllers.ProfileUser))).Methods("GET")
	r.HandleFunc("/profile", middleware.MiddleWare(http.HandlerFunc(profileControllers.UpdateProfile))).Methods("PATCH")
}
