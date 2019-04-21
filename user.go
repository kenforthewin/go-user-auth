package main

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Password  string    `json:"password,omitempty" db:"password"`
}

type Users []User
