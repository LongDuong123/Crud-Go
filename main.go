package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
