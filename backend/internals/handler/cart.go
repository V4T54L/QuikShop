package handler

import (
	"backend/internals/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) GetUserCartHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get userId from token
	userID := 1
	cart, err := h.cartService.GetUserCart(r.Context(), userID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	productIDs := make([]int, 0, len(cart))
	for _, item := range cart {
		productIDs = append(productIDs, item.ProductID)
	}
	products, err := h.productService.GetProductByIDs(r.Context(), productIDs)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cartDetail := make([]models.CartItemDetail, 0, len(products))
	for _, product := range products {
		item := models.CartItemDetail{
			ProductID:    product.ID,
			ProductName:  product.Name,
			CurrentPrice: float32(product.Price / 100),
		}

		for _, cartItem := range cart {
			if cartItem.ProductID == product.ID {
				item.Quantity = cartItem.Quantity
				break
			}
		}
	}

	jsonResponse(w, http.StatusOK, cartDetail)
}

func (h *Handler) UpdateCartItemHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.CartItem
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if payload.Quantity == 0 {
		errorResponse(w, http.StatusBadRequest, "quantity can't be negative")
		return
	}

	// TODO: get UserID from token
	userID := 1

	var cart *models.CartItem
	if payload.Quantity > 0 {
		cart, err = h.cartService.AddItemToCart(r.Context(), userID, payload.ProductID, payload.Quantity)
	} else {
		cart, err = h.cartService.RemoveItemFromCart(r.Context(), userID, payload.ProductID, -payload.Quantity)
	}
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, cart)
}
