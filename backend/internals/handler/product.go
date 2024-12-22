package handler

import (
	"backend/internals/models"
	"backend/internals/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	products, err := ph.service.GetProducts(ctx)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, products)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.ProductDetail{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ph.service.CreateProduct(r.Context(), product); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusCreated, product)
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.ProductDetail{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ph.service.UpdateProduct(r.Context(), product); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, product)
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("productID")
	if productID == "" {
		errorResponse(w, http.StatusBadRequest, "product ID is required")
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	if err := ph.service.DeleteProduct(r.Context(), id); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"message": "product deleted"})
}
