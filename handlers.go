package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
	return
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r)
	if user.ID > 0 {
		user.Password = ""
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	u := RepoInsertUser(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

type SigninResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expiration"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := RepoFindUserByEmail(creds.Email)

	if u.ID < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = CompareHash(creds.Password, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(8760 * time.Hour)
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := SigninResponse{
		Token:   tokenString,
		Expires: expirationTime.String(),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
	return
}
