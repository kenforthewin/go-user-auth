package main

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY_BASE"))

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateHash(plainText string) string {
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(passwordHashInBytes)
}

func CompareHash(plainText string, hashText string) error {
	plainTextInBytes := []byte(plainText)
	hashTextInBytes := []byte(hashText)
	return bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
}
