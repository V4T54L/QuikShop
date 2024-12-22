package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"context"
	"errors"
)

type ProductService struct {
	store store.ProductStore
}

func NewProductService(store store.ProductStore) *ProductService {
	return &ProductService{store: store}
}

func (ps *ProductService) GetProducts(ctx context.Context) ([]models.ProductDetail, error) {
	return ps.store.GetProducts(ctx)
}

func (ps *ProductService) CreateProduct(ctx context.Context, product *models.ProductDetail) error {
	// validate product details
	if product.Name == "" || product.Price == 0 || product.Stock == 0 {
		return errors.New("name, price and stock are required")
	}

	return ps.store.CreateProduct(ctx, product)
}

func (ps *ProductService) UpdateProduct(ctx context.Context, product *models.ProductDetail) error {
	if product.ID == 0 {
		return errors.New("product ID is required")
	}
	if product.Name == "" || product.Price == 0 || product.Stock == 0 {
		return errors.New("name, price and stock are required")
	}

	return ps.store.UpdateProduct(ctx, product)
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("product ID is required")
	}

	return ps.store.DeleteProduct(ctx, id)
}
