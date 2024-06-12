package routes

import (
	controllers "crud/Api/Controllers"
	usecase "crud/Application/UseCase"
	domain "crud/Domain"

	"github.com/gorilla/mux"
)

func RegisterRoutesSignUp(r *mux.Router, ur domain.UserRepository) {
	SignupUseCase := usecase.NewSignUpUseCase(ur)
	SignupControllers := controllers.NewSignUpControllers(SignupUseCase)
	r.HandleFunc("/signup", SignupControllers.SignUp).Methods("POST")
}
