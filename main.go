package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"))
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
