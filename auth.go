package main

import (
	// "net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func createJWTKey(userName string, expirationTime time.Time, Role string) (string, error) {
	claim := jwt.MapClaims{
		"user_Name": userName,
		"role":      Role,
		"exp":       expirationTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// func VerifyToken(cookie *http.Cookie){
// 	tokenCookie := cookie.Value

// 	tkn, err := jwt.Parse(tokenCookie, func(t *jwt.Token) (interface{},error) {
// 		return jwtKey, nil
// 	})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {

// 		}
// 	}

// }
