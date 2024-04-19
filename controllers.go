package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func loginUser(w http.ResponseWriter, r *http.Request) {

	var credentitals User

	err := json.NewDecoder(r.Body).Decode(&credentitals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	checkPassWord, err := db.Query("SELECT password, role FROM users WHERE email = ?", credentitals.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var check string
	if checkPassWord.Next() {
		err := checkPassWord.Scan(&check, &credentitals.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err1 := bcrypt.CompareHashAndPassword([]byte(check), []byte(credentitals.Password))
	if err1 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Incorrect password")
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	accessToken, err := createJWTKey(credentitals.Email, expirationTime, credentitals.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   accessToken,
		Expires: expirationTime,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()

	for _, cookie := range cookies {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logout Successful")
}

func checkEmailUser(user User) (int, error) {
	if len(user.Email) <= 10 {
		return 0, nil
	}
	check_Email, err := db.Query("SELECT email FROM users WHERE email = ?", user.Email)
	if err != nil {
		return 0, err
	}
	defer check_Email.Close()
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
	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	resultCheckEmail, err := checkEmailUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resultCheckEmail == 0 {
		http.Error(w, "Email phải lớn hơn 10 ký tự", http.StatusBadRequest)
		return
	} else if resultCheckEmail == 1 {
		http.Error(w, "Email đã tồn tại", http.StatusBadRequest)
		return
	}

	resultCheckPassWord := checkPassWord(user)

	if resultCheckPassWord == 0 {
		http.Error(w, "Password phải có hơn 8 ký tự với ít nhất một số , một từ và một ký tự đặc biệt", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("INSERT INTO users (name, age, gender, email, password) VALUES (?, ?, ?, ?, ?)", user.Name, user.Age, user.Gender, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	var update_User User
	err := json.NewDecoder(r.Body).Decode(&update_User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(update_User.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err1 := db.Exec("UPDATE users SET name = ?, age = ?, gender = ?, password = ? WHERE id = ?", update_User.Name, update_User.Age, update_User.Gender, hashPassWord, userID)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Update Successful")
}

func getProductAll(w http.ResponseWriter, r *http.Request) {

	getDatabaseProduct, err := db.Query("SELECT * FROM product")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var listProduct []Product

	for getDatabaseProduct.Next() {
		var get_Product Product
		err := getDatabaseProduct.Scan(&get_Product.ID, &get_Product.Name, &get_Product.Image_url, &get_Product.Price, &get_Product.Create_By)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		listProduct = append(listProduct, get_Product)
	}
	jsonData, _ := json.MarshalIndent(listProduct, "", "  ")
	w.Write(jsonData)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	pageSizeStr := r.URL.Query().Get("page_size")
	pageNumberStr := r.URL.Query().Get("page_number")

	if pageSizeStr == "" || pageNumberStr == "" {
		http.Error(w, "Missing numbers", http.StatusBadRequest)
		return
	}

	pageSize, _ := strconv.Atoi(pageSizeStr)
	pageNumber, _ := strconv.Atoi(pageNumberStr)

	startPage := (pageSize - 1) * pageNumber

	getDataProductPage, err := db.Query("SELECT * FROM product WHERE id >= $1 AND id <= $2", startPage, startPage+pageNumber)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var getProductPage []Product
	for getDataProductPage.Next() {
		var product Product
		err := getDataProductPage.Scan(&product.ID, &product.Name, &product.Image_url, &product.Price, &product.Create_By)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		getProductPage = append(getProductPage, product)
	}
	jsonData, _ := json.MarshalIndent(getProductPage, "", "  ")
	w.Write(jsonData)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tonkenStr := cookie.Value

	tkn, err := jwt.Parse(tonkenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := tkn.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		fmt.Fprintln(w, "You are not admin")
		return
	}
	vars := mux.Vars(r)
	productID := vars["id"]

	var update_Product Product

	err1 := json.NewDecoder(r.Body).Decode(&update_Product)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	_, err2 := db.Exec("UPDATE product SET name = $1 , image_url = $2 , price = ? WHERE id = ?", update_Product.Name, update_Product.Image_url, update_Product.Price, productID)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Update Successful")
}
