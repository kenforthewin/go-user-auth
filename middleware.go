package main

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := strings.Split(r.Header.Get("Authorization"), " ")

		if len(header) < 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if header[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := header[1]
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if user := RepoFindUserByEmail(claims.Email); user.ID > 0 {
			SetUser(r, user)
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	})
}

func SetUser(r *http.Request, val User) {
	context.Set(r, 0, val)
}

func GetUser(r *http.Request) User {
	if rv := context.Get(r, 0); rv != nil {
		return rv.(User)
	}
	return User{}
}
