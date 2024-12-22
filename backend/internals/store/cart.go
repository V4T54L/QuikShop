package store

import (
	"backend/internals/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type CartStore interface {
	GetCart(ctx context.Context, userID int) (*models.CartDetail, error)
	AddToCart(ctx context.Context, userID int, item *models.CartItem) error
	RemoveFromCart(ctx context.Context, userID, productID int) error
	ClearCart(ctx context.Context, userID int) error
}

type CartStoreImpl struct {
	db *sql.DB
}

func NewCartStore(db *sql.DB) *CartStoreImpl {
	return &CartStoreImpl{db: db}
}

func (cs *CartStoreImpl) GetCart(ctx context.Context, userID int) (*models.CartDetail, error) {
	var cartDetail models.CartDetail
	var cartItems []models.CartItem

	// Execute the query
	rows, err := cs.db.QueryContext(ctx, "SELECT c.CartID AS id, ci.CartItemID AS cart_item_id, ci.ProductID, ci.Quantity FROM Cart c JOIN CartItems ci ON ci.CartID = c.CartID WHERE c.UserID = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(&cartDetail.ID, &item.ID, &item.ProductID, &item.Quantity)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}

	cartDetail.Items = cartItems

	return &cartDetail, nil
}

func (cs *CartStoreImpl) AddToCart(ctx context.Context, userID int, item *models.CartItem) error {
	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var cartID int
	err = tx.QueryRowContext(ctx, "SELECT CartID FROM Cart WHERE UserID = $1", userID).Scan(&cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("cart not found for user ID %d", userID)
		}
		return err
	}

	_, err = tx.Exec("INSERT INTO CartItems (CartID, ProductID, Quantity) VALUES ($1, $2, $3)", cartID, item.ProductID, item.Quantity)
	if err != nil {
		if isUniqueViolationError(err) {
			_, updateErr := tx.Exec("UPDATE CartItems SET Quantity = Quantity + $1 WHERE CartID = $2 AND ProductID = $3", item.Quantity, cartID, item.ProductID)
			if updateErr != nil {
				return updateErr
			}
		} else {
			return err
		}
	}

	return tx.Commit()
}

func isUniqueViolationError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "unique constraint")
}

func (cs *CartStoreImpl) RemoveFromCart(ctx context.Context, userID, productID int) error {
	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var cartID int
	err = tx.QueryRowContext(ctx, "SELECT CartID FROM Cart WHERE UserID = $1", userID).Scan(&cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("cart not found for user ID %d", userID)
		}
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM CartItems WHERE CartID = $1 AND ProductID = $2", cartID, productID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (cs *CartStoreImpl) ClearCart(ctx context.Context, userID int) error {
	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var cartID int
	err = tx.QueryRowContext(ctx, "SELECT CartID FROM Cart WHERE UserID = $1", userID).Scan(&cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("cart not found for user ID %d", userID)
		}
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM CartItems WHERE CartID = $1", cartID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
