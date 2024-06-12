package internal

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func CreateToken(id int, userName string, expirationTime time.Time) (string, error) {
	claim := jwt.MapClaims{
		"Id":        id,
		"user_Name": userName,
		"exp":       expirationTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyTokenAndGetID(tokenString string) (int, error) {
	tkn, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return -1, err
	}

	if !tkn.Valid {
		return -1, err
	}
	claim := tkn.Claims.(jwt.MapClaims)
	idFloat64 := claim["Id"].(float64)

	return int(idFloat64), nil
}
