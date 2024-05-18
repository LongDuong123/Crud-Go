package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	if err := Connnect(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	RegisterRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "DELETE", "POST"}),
		handlers.AllowedHeaders([]string{"Authorization", "Context-Type"}),
	)(r))
}
