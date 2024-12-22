package store

import (
	"backend/internals/models"
	"context"
	"database/sql"
	"fmt"
)

type ProductStore interface {
	GetProducts(ctx context.Context) ([]models.ProductDetail, error)
	CreateProduct(ctx context.Context, product *models.ProductDetail) error
	UpdateProduct(ctx context.Context, product *models.ProductDetail) error
	DeleteProduct(ctx context.Context, id int) error
}

type ProductStoreImpl struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStoreImpl {
	return &ProductStoreImpl{db: db}
}

func (ps *ProductStoreImpl) GetProducts(ctx context.Context) ([]models.ProductDetail, error) {
	stmt, err := ps.db.PrepareContext(ctx, `SELECT ProductID, Name, Description, Price, Stock, ImageURI from Products`)
	if err != nil {
		return nil, fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not execute statement: %w", err)
	}
	defer rows.Close()

	var products []models.ProductDetail
	for rows.Next() {
		var product models.ProductDetail
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.ImageURI); err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// stmt, err := ps.db.PrepareContext(ctx, `INSERT INTO Products (Name, Description, Price, Stock, ImageURI) VALUES ($1, $2, $3, $4, $5)`)
func (ps *ProductStoreImpl) CreateProduct(ctx context.Context, product *models.ProductDetail) error {

	tx, err := ps.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO Products (Name, Description, Price, Stock, ImageURI) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, product.Name, product.Description, product.Price, product.Stock, product.ImageURI).Scan(&product.ID); err != nil {
		return fmt.Errorf("could not execute statement: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	fmt.Printf("Product created with ID: %d\n", product.ID)
	return nil
}

func (ps *ProductStoreImpl) UpdateProduct(ctx context.Context, product *models.ProductDetail) error {
	tx, err := ps.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.PrepareContext(ctx, `UPDATE Products SET Name=$1, Description=$2, Price=$3, Stock=$4, ImageURI=$5 WHERE ProductID=$6`)
	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, product.Name, product.Description, product.Price, product.Stock, product.ImageURI, product.ID); err != nil {
		return fmt.Errorf("could not execute statement: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	fmt.Printf("Product updated with ID: %d\n", product.ID)
	return nil
}

func (ps *ProductStoreImpl) DeleteProduct(ctx context.Context, id int) error {
	tx, err := ps.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.PrepareContext(ctx, `DELETE FROM Products WHERE ProductID=$1`)
	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("could not execute statement: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	fmt.Printf("Product deleted with ID: %d\n", id)
	return nil
}
