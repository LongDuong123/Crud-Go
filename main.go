package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO user (Name,Age,Email,Password) VALUES (? , ? , ? , ?)", user.Name, user.Age, user.Email, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println("Create Successful")
}

func readUser(w http.ResponseWriter, r *http.Request) {
	getUser, err := db.Query("SELECT * FROM user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []User
	for getUser.Next() {
		var user User
		err := getUser.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE user SET Name = ? , Age = ? , Email = ? , Password = ? WHERE ID = ?", user.Name, user.Age, user.Email, user.Password, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Update Successful")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	_, err := db.Exec("DELETE FROM user WHERE ID = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Delete Successful")
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:admin123@/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", readUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
