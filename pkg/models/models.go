package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

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
