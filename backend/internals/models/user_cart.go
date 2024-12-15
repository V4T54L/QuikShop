package models

type CartItemDetail struct {
	ProductID    int     `json:"id"`
	ProductName  string  `json:"name"`
	Quantity     int     `json:"quantity"`
	CurrentPrice float32 `json:"current_price"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
