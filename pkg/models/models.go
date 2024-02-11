package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Article struct {
	ID          int       `json:"id"`
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	Readership  string    `json:"readership"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"created"`
	Content     string    `json:"content"`
}
type Department struct {
	ID            int    `json:"id"`
	DepName       string `json:"dep_name"`
	StaffQuantity int    `json:"staff_quantity"`
}
type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Approved bool   `json:"approved"`
}
