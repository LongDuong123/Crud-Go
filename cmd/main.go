package main

import (
	routes "crud/Api/Routes"
	"crud/config"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Init, err := config.InitializeAppConfig()
	if err != nil {
		log.Fatal("Fail init")
	}
	r := mux.NewRouter()
	routes.RegisterRoutesLogin(r, Init.RepositoryUser)
	routes.RegisterRoutesProfile(r, Init.RepositoryUser)
	routes.RegisterRoutesSignUp(r, Init.RepositoryUser)
	routes.RegisterRoutesProduct(r, Init.RepositoryProduct, Init.RepositoryProductRedis)
	http.Handle("/", r)
	http.ListenAndServe(":8080", (r))
}
