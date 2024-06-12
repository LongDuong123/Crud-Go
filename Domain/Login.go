package domain

import "time"

type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
type LoginResponse struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}

type LoginInteractor interface {
	GetByEmail(email string) (*User, error)
	CheckPassword(hashpassword string, password string) bool
	CreateAccessToken(id int, userName string, expiry time.Time) (string, error)
	CreateRefreshToken(id int, userName string, expiry time.Time) (string, error)
}
