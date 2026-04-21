package main

import "time"

type Product struct {
	ID          int64     `json:"id"`
	ProductCode string    `json:"product_code" binding:"required,min=3,max=20"`
	Name        string    `json:"name" binding:"required,min=2,max=100"`
	Price       float64   `json:"price" binding:"required,gt=0"`
	Quantity    int       `json:"quantity" binding:"gte=0"`
	Category    string    `json:"category" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
