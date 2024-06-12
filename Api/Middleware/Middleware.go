package middleware

import (
	"context"
	"crud/internal"
	"net/http"
)

type contextKey string

const UserIdKey contextKey = "user-id"

func MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("AccessToken")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userID, err1 := internal.VerifyTokenAndGetID(cookies.Value)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), UserIdKey, userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
