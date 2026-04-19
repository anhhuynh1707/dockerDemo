package main

type Product struct {
	ID          int64   `json:"id"`
	ProductCode string  `json:"product_code"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
}
