package models

import "time"

type Product struct {
	ID         int       `json:"id"`
	CategoryId int       `json:"category_id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProductDetail struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	Category  Category  `json:"category"`
}
