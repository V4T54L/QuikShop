package models

type CartItemDetail struct {
	ProductID    string  `json:"id"`
	ProductName  string  `json:"name"`
	Quantity     int     `json:"quantity"`
	CurrentPrice float32 `json:"current_price"`
}

type CartItem struct {
	ProductID int
	Quantity  int
}
