package store

import (
	"backend/internals/models"
	"context"
)

type ProductStore interface {
	GetProductByID(ctx context.Context, productID int) (*models.ProductDetail, error)
	SearchProducts(ctx context.Context, query string, start int, limit int) ([]models.ProductSummary, error)
}
