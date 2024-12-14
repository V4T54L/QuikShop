package store

import (
	"backend/internals/models"
	"context"
)

type UserCartStore interface {
	AddItemToUserCart(ctx context.Context, userID int, productID int, quantity int) (*models.CartItem, error)
	GetUserCart(ctx context.Context, userID int) ([]models.CartItem, error)
	RemoveItemFromUserCart(ctx context.Context, userID int, productID int, quantity int) (*models.CartItem, error)
}
