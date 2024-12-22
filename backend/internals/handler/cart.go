package handler

import (
	"backend/internals/models"
	"backend/internals/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type CartHandler struct {
	service services.CartService
}

func NewCartHandler(service services.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (ch *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	// get user ID from request context
	userID := r.Context().Value("userID").(int)

	cart, err := ch.service.GetCart(r.Context(), userID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, cart)
}

func (ch *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	// get user ID from request context
	userID := r.Context().Value("userID").(int)

	// decode request body
	itemDetail := &models.CartItem{}
	err := json.NewDecoder(r.Body).Decode(&itemDetail)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// add item to cart
	if err := ch.service.AddToCart(r.Context(), userID, itemDetail); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusCreated, itemDetail)
}

func (ch *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	// get user ID from request context
	userID := r.Context().Value("userID").(int)

	// get product ID from request URL
	productID, err := strconv.Atoi(r.URL.Query().Get("product_id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	// remove item from cart
	if err := ch.service.RemoveFromCart(r.Context(), userID, productID); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, nil)
}

func (ch *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	// get user ID from request context
	userID := r.Context().Value("userID").(int)

	// clear cart
	if err := ch.service.ClearCart(r.Context(), userID); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, nil)
}
