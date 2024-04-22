package main

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", createUser).Methods("POST")
	r.HandleFunc("/login", loginUser).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/productAll", getProductAll).Methods("GET")
	r.HandleFunc("/product", getProduct).Methods("GET")
	r.HandleFunc("/product/{id}", VerifyToken(updateProduct)).Methods("PUT")
	r.HandleFunc("/product/{id}", VerifyToken(deleteProduct)).Methods("DELETE")
}
