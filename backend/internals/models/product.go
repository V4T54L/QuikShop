package models

type ProductDetail struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageURI    string `json:"images"`
	Stock       int    `json:"stock"`
}
