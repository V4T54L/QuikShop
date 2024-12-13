package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"context"
	"time"
)

const (
	maxQueryTime = 5 * time.Second
)

type ProductService struct {
	store store.ProductStore
}

func NewProductService(store store.ProductStore) *ProductService {
	return &ProductService{store: store}
}

func (s *ProductService) GetProductsBySearch(mainCtx context.Context, query string, start, limit int) ([]models.ProductSummary, error) {
	ctx, cancel := context.WithTimeout(mainCtx, maxQueryTime)
	defer cancel()
	return s.store.SearchProducts(ctx, query, start, limit)
}

func (s *ProductService) GetProductDetail(mainCtx context.Context, productID int) (*models.ProductDetail, error) {
	ctx, cancel := context.WithTimeout(mainCtx, maxQueryTime)
	defer cancel()
	return s.store.GetProductByID(ctx, productID)
}
