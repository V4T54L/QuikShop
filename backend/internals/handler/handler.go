package handler

import "backend/internals/services"

type Handler struct {
	productService services.ProductService
	cartService    services.CartService
}

func NewHandler(productService services.ProductService, cartService services.CartService) *Handler {
	return &Handler{productService: productService, cartService: cartService}
}
