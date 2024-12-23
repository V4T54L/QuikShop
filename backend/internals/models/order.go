package models

type OrderItem struct {
	ID        int `json:"id"`
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
}

type OrderDetail struct {
	ID        int         `json:"id"`
	UserID    int         `json:"userID"`
	Items     []OrderItem `json:"items"`
	Total     int         `json:"total"`
	Status    string      `json:"status"`
	CreatedAt string      `json:"created_at"`
}

type OrderStatus struct {
	Status string `json:"status"`
}
