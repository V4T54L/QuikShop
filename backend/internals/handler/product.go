package handler

import (
	"backend/internals/services"
	"log"
	"net/http"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) SearchProductHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	log.Println("Query : ", query)
}
