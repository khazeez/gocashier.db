package models

import "time"

type Product struct {
	ID         int       `json:"id"`
	CategoryId int       `json:"category_id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProductDetail struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Stock    int      `json:"stock"`
	Category Category `json:"category"`
}
