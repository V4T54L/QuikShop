package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"context"
)

type OrderService struct {
	store store.OrderStore
}

func NewOrderService(store store.OrderStore) *OrderService {
	return &OrderService{store: store}
}

func (os *OrderService) CreateOrder(ctx context.Context, order *models.OrderDetail) error {
	return os.store.CreateOrder(ctx, order)
}

func (os *OrderService) GetOrders(ctx context.Context, userID int) ([]models.OrderDetail, error) {
	return os.store.GetOrders(ctx, userID)
}

func (os *OrderService) GetOrder(ctx context.Context, orderID int) (*models.OrderDetail, error) {
	return os.store.GetOrder(ctx, orderID)
}

func (os *OrderService) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	return os.store.UpdateOrderStatus(ctx, orderID, status)
}
