package store

import (
	"backend/internals/models"
	"context"
)

type ProductStore interface {
	GetProductByID(ctx context.Context, productID int) (*models.ProductDetail, error)
	GetProductsByIDs(ctx context.Context, ids []int) ([]models.ProductSummary, error)
	SearchProducts(ctx context.Context, query string, start int, limit int) ([]models.ProductSummary, error)
}
