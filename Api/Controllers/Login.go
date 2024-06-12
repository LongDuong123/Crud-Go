package controllers

import (
	domain "crud/Domain"
	"encoding/json"
	"net/http"
	"time"
)

type LoginController struct {
	LoginUseCase domain.LoginInteractor
}

func NewLoginController(lu domain.LoginInteractor) *LoginController {
	return &LoginController{
		LoginUseCase: lu,
	}
}

func (lc *LoginController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var UserLogin domain.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&UserLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	User, err := lc.LoginUseCase.GetByEmail(UserLogin.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	ok := lc.LoginUseCase.CheckPassword(UserLogin.Password, User.Password)
	if !ok {
		http.Error(w, "Wrong Password", http.StatusBadRequest)
		return
	}
	accessToken, err := lc.LoginUseCase.CreateAccessToken(User.ID, User.Name, time.Now().Add(1*time.Hour))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := lc.LoginUseCase.CreateRefreshToken(User.ID, User.Name, time.Now().Add(24*time.Hour))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookies := []http.Cookie{
		{
			Name:    "AccessToken",
			Value:   accessToken,
			Expires: time.Now().Add(1 * time.Hour),
		},
		{
			Name:    "RefreshToken",
			Value:   refreshToken,
			Expires: time.Now().Add(24 * time.Hour),
		},
	}
	for i := 0; i < len(cookies); i++ {
		http.SetCookie(w, &cookies[i])
	}
}
