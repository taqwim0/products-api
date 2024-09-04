package model

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Variety     string  `json:"variety"`
	Rating      float64 `json:"rating"`
	Stock       int     `json:"stock"`
}
