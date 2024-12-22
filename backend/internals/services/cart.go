package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"context"
)

type CartService struct {
	store store.CartStore
}

func NewCartService(store store.CartStore) *CartService {
	return &CartService{store: store}
}

func (cs *CartService) GetCart(ctx context.Context, userID int) (*models.CartDetail, error) {
	return cs.store.GetCart(ctx, userID)
}

func (cs *CartService) AddToCart(ctx context.Context, userID int, item *models.CartItem) error {
	return cs.store.AddToCart(ctx, userID, item)
}

func (cs *CartService) RemoveFromCart(ctx context.Context, userID, productID int) error {
	return cs.store.RemoveFromCart(ctx, userID, productID)
}

func (cs *CartService) ClearCart(ctx context.Context, userID int) error {
	return cs.store.ClearCart(ctx, userID)
}
