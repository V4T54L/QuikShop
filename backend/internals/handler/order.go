package handler

import (
	"backend/internals/models"
	"backend/internals/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	orderDetail := &models.OrderDetail{}
	if err := json.NewDecoder(r.Body).Decode(orderDetail); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := oh.orderService.CreateOrder(r.Context(), orderDetail); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusCreated, orderDetail)
}

func (oh *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	orders, err := oh.orderService.GetOrders(r.Context(), userID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, orders)
}

func (oh *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(r.URL.Query().Get("order_id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid order ID")
		return
	}

	order, err := oh.orderService.GetOrder(r.Context(), orderID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, order)
}

func (oh *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(r.URL.Query().Get("order_id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid order ID")
		return
	}

	status := &models.OrderStatus{}
	if err := json.NewDecoder(r.Body).Decode(status); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := oh.orderService.UpdateOrderStatus(r.Context(), orderID, status.Status); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, status)
}
