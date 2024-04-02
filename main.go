package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtKey = []byte("secret_key")

type User struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	//Gender   string   `json:"Gender"`
	Email string `json:"Email"`
	//Cart     []string `json:"Cart"`
	//Role     string `json:"Role"`
	Password string `json:"Password"`
}

type Product struct {
	ID             int
	Name           string
	Image          string
	CreateDate     time.Time
	UpdateDate     time.Time
	IsCreatedBy    int
	ThumbnailImage string
}

type Credentials struct {
	Username string 	`json:"username"`
 	Password string		`json:"password"`
}
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func createToken(userID int) (string,  error) {

	claims := &jwt.MapClaims{
		"user_id": userID,
		"Expires": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString , nil
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.Decoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	if credentials == "" {

	}

	checkUserName , err := db.Query("SELECT Password FROM user WHERE Email = ?",credentials.Username)

	
}

func checkEmailUser(user User) (int, error) {
	if len(user.Email) <= 10 {
		return 0, nil
	}
	check_Email, err := db.Query("SELECT Email FROM user WHERE Email = ?", user.Email)
	if err != nil {
		return 0, err
	}
	var check string
	if check_Email.Next() {
		err = check_Email.Scan(&check)
		if err != nil {
			return 0, err
		}
	}
	if check != "" {
		return 1, nil
	}
	return 2, nil
}

func checkPassWord(user User) int {
	if len(user.Password) < 8 {
		return 0
	}
	check := []int{0, 0, 0}
	for _, k := range user.Password {
		if k >= 48 && k <= 57 {
			check[0]++
		} else if k >= 65 && k <= 90 || k >= 97 && k <= 122 {
			check[1]++
		} else if k >= 33 && k <= 126 {
			check[2]++
		}
	}
	for _, k := range check {
		if k == 0 {
			return 0
		}
	}
	return 1
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultCheckEmail, err := checkEmailUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resultCheckEmail == 0 {
		fmt.Fprint(w, "Email phải lớn hơn 10 ký tự")
		return
	} else if resultCheckEmail == 1 {
		fmt.Fprint(w, "Email đã tồn tại")
		return
	}

	resultCheckPassWord := checkPassWord(user)

	if resultCheckPassWord == 0 {
		fmt.Fprint(w, "Password phải có hơn 8 ký tự với ít nhất một số , một từ và một ký tự đặc biệt")
		return
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	_, err = db.Exec("INSERT INTO user (Name,Age,Email,Password) VALUES (? , ? , ? , ?)", user.Name, user.Age, user.Email, newPassword)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Lỗi khi hash mật khẩu:", err)
		return
	}

	_, err = db.Exec("UPDATE user SET Name = ? , Age = ? , Email = ? , Password = ? WHERE ID = ?", user.Name, user.Age, user.Email, newPassword, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	_, err := db.Exec("DELETE FROM user WHERE ID = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err1 := db.Exec("INSERT INTO product (Name,Image,Createdate,UpdateDate,IsCreatedBy,ThumbnailImage) VALUE (?,?,?,?,?,?)", product.Name, product.Image, product.CreateDate, product.UpdateDate, product.IsCreatedBy, product.ThumbnailImage)

	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func readProduct(w http.ResponseWriter, r *http.Request) {
	_, err := db.Query("SELECT * FROM product ")
	if err 
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

	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products", readProduct).Methods("GET")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	r.HandleFunc("/login" , loginUser)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
