package store

import (
	"backend/internals/models"
	"context"
	"errors"
)

type MockUserCartStore struct {
	products []models.CartItem
}

func NewMockCartStore() *MockUserCartStore {
	return &MockUserCartStore{products: []models.CartItem{}}
}

func (s *MockUserCartStore) GetUserCart(ctx context.Context, userID int) ([]models.CartItem, error) {
	products := make([]models.CartItem, 0, len(s.products))
	for _, v := range products {
		if v.Quantity == 0 {
			continue
		}
		products = append(products, v)
	}
	return s.products, nil
}

func (s *MockUserCartStore) AddItemToUserCart(ctx context.Context, userID, productID, quantity int) (*models.CartItem, error) {
	var item *models.CartItem
	for _, v := range s.products {
		if v.ProductID == productID {
			item = &v
			break
		}
	}
	if item == nil {
		item = &models.CartItem{ProductID: productID, Quantity: quantity}
	} else {
		item.Quantity += quantity
	}

	s.products = append(s.products, *item)
	return item, nil
}

func (s *MockUserCartStore) RemoveItemFromUserCart(ctx context.Context, userID, productID, quantity int) (*models.CartItem, error) {
	var item *models.CartItem
	for _, v := range s.products {
		if v.ProductID == productID {
			item = &v
			break
		}
	}
	if item == nil {
		return nil, errors.New("product does not exist in user's cart")
	} else {
		item.Quantity -= quantity
	}
	if item.Quantity < 0 {
		item.Quantity = 0
	}

	return nil, nil
}
