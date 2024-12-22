package handler

import "backend/internals/services"

type Handler struct {
	productService services.ProductService
	cartService    services.CartService
	userService    services.UserService
}

func NewHandler(productService services.ProductService, cartService services.CartService, userService services.UserService) *Handler {
	return &Handler{productService: productService, cartService: cartService, userService: userService}
}
