package controllers

import (
	domain "crud/Domain"
	"encoding/json"
	"net/http"
)

type SignUpControllers struct {
	signUpUseCase domain.SignupInteractor
}

func NewSignUpControllers(signUp domain.SignupInteractor) *SignUpControllers {
	return &SignUpControllers{signUpUseCase: signUp}
}

func (sg *SignUpControllers) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUp domain.Signup
	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err := sg.signUpUseCase.CheckUserByEmail(signUp.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Email has been used", http.StatusBadRequest)
		return
	}
	err = sg.signUpUseCase.Create(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
