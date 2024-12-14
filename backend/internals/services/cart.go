package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"context"
)

type CartService struct {
	store store.UserCartStore
}

func NewCartService(store store.UserCartStore) *CartService {
	return &CartService{store: store}
}

func (s *CartService) GetUserCart(mainCtx context.Context, userID int) ([]models.CartItem, error) {
	ctx, cancel := context.WithTimeout(mainCtx, maxQueryTime)
	defer cancel()
	return s.store.GetUserCart(ctx, userID)
}

func (s *CartService) AddItemToCart(mainCtx context.Context, userID, productID, quantity int) (*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(mainCtx, maxQueryTime)
	defer cancel()
	return s.store.AddItemToUserCart(ctx, userID, productID, quantity)
}

func (s *CartService) RemoveItemFromCart(mainCtx context.Context, userID, productID, quantity int) (*models.CartItem, error) {
	ctx, cancel := context.WithTimeout(mainCtx, maxQueryTime)
	defer cancel()
	return s.store.RemoveItemFromUserCart(ctx, userID, productID, quantity)
}
