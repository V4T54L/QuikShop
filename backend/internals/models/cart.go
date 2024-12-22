package models

type CartItem struct {
	ID        int `json:"id"`
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type CartDetail struct {
	ID    int        `json:"id"`
	Items []CartItem `json:"items"`
}
