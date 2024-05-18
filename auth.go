package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func createTokenAndSetCookie(userName string, expirationTime time.Time, Role string, typeToken string, w http.ResponseWriter) error {
	claim := jwt.MapClaims{
		"user_Name": userName,
		"role":      Role,
		"exp":       expirationTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:    typeToken,
		Value:   tokenString,
		Expires: expirationTime,
	}

	http.SetCookie(w, &cookie)
	return nil
}

func MiddleWare(next http.HandlerFunc, nextNoAuth http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err1 := verifyToken(r, "AccessToken")
		refreshToken, err2 := verifyToken(r, "RefreshToken")
		if !err1 && !err2 {
			nextNoAuth.ServeHTTP(w, r)
			return
		} else if !err1 && err2 {
			claim := refreshToken.Claims.(jwt.MapClaims)
			err := createTokenAndSetCookie(claim["user_Name"].(string), time.Now().Add(time.Hour*1), claim["role"].(string), "AccessToken", w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = createTokenAndSetCookie(claim["user_Name"].(string), time.Now().Add(time.Hour*24), claim["role"].(string), "RefreshToken", w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if claim["role"] != "admin" {
				return
			}
			next.ServeHTTP(w, r)
			return
		}
		Claims := accessToken.Claims.(jwt.MapClaims)
		if Claims["role"] != "admin" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func verifyToken(r *http.Request, nameToken string) (*jwt.Token, bool) {
	token, err := r.Cookie(nameToken)
	if err != nil {
		return nil, false
	}

	tokenString := token.Value

	tkn, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, false
	}

	if !tkn.Valid {
		return nil, false
	}
	return tkn, true
}
