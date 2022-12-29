package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task_2/model"
	"task_2/service"
	"time"
)

var storage *service.MemoryStorage

func UserRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// fmt.Fprintf(w, storage.Users)
	}
	if r.Method == "POST" {
		var u model.User
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Fprintf(w, storage.CreateUser(u.Name, u.Login, u.Password, u.Phone, u.BirthDate))
	}
	if r.Method == "PUT" {
		var u model.User
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Fprintf(w, storage.UpdateUser(u.Id, u.Name, u.Login, u.Password, u.Phone, u.BirthDate))
	}
}

func UserLoginRequest(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	fmt.Fprintf(w, "Login user")
}
func main() {
	storage = service.NewMemoryStorage()
	mux := http.NewServeMux()
	mux.HandleFunc("/user", UserRequest)
	mux.HandleFunc("/user/login", UserLoginRequest)
	//We can't use 80 port without root permissions
	http.ListenAndServe(":8080", mux)
}
