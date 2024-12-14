package handler

import (
	"backend/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) SearchProductHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	pageNo, err := utils.GetInt(r.URL.Query().Get("pageNo"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if pageNo == nil {
		pgNo := 0
		pageNo = &pgNo
	}
	limit, err := utils.GetInt(r.URL.Query().Get("limit"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if limit == nil {
		lim := 10
		limit = &lim
	}

	products, err := h.productService.GetProductsBySearch(r.Context(), query, *pageNo, *limit)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, products)
}

func (h *Handler) GetProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetInt(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// id=nil ? not handled by search product handler
	// if id == nil {
	// 	errorResponse(w, http.StatusBadRequest, "id not provided in path param")
	// 	return
	// }

	productDetail, err := h.productService.GetProductDetail(r.Context(), *id)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, productDetail)
}