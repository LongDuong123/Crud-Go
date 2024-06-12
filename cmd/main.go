package main

import (
	routes "crud/Api/Routes"
	mysql "crud/Infrastructure/database/MySQL"
	redis "crud/Infrastructure/database/Redis"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	databaseMySql, err := mysql.ConnnectMySql()
	if err != nil {
		log.Fatal(err)
	}
	databaseRedis := redis.ConnnectRedis()
	repositoryUser := mysql.NewUserRepository(databaseMySql)
	repositoryProduct := mysql.NewProductRepository(databaseMySql)
	repositoryProductRedis := redis.NewProductRepositoryByRedis(databaseRedis)
	r := mux.NewRouter()
	routes.RegisterRoutesLogin(r, repositoryUser)
	routes.RegisterRoutesProfile(r, repositoryUser)
	routes.RegisterRoutesSignUp(r, repositoryUser)
	routes.RegisterRoutesProduct(r, repositoryProduct, repositoryProductRedis)
	http.Handle("/", r)
	http.ListenAndServe(":8080", (r))
}
