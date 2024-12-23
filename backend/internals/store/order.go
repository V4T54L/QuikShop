package store

import (
	"backend/internals/models"
	"context"
	"database/sql"
)

type OrderStore interface {
	CreateOrder(ctx context.Context, order *models.OrderDetail) error
	GetOrders(ctx context.Context, userID int) ([]models.OrderDetail, error)
	GetOrder(ctx context.Context, orderID int) (*models.OrderDetail, error)
	// UpdateOrderStatus(ctx, orderID, status)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
}

type OrderStoreImpl struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStoreImpl {
	return &OrderStoreImpl{db: db}
}

func (os *OrderStoreImpl) CreateOrder(ctx context.Context, order *models.OrderDetail) error {
	tx, err := os.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.PrepareContext(ctx, `
        INSERT INTO Orders (UserID, TotalAmount, Status)
        VALUES ($1, $2, $3) RETURNING OrderID
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, order.UserID, order.Total, order.Status).Scan(&order.ID); err != nil {
		return err
	}

	for _, item := range order.Items {
		stmt, err := tx.PrepareContext(ctx, `
            INSERT INTO OrderItems (OrderID, ProductID, Quantity, Price)
            VALUES ($1, $2, $3, $4)
        `)
		if err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx, order.ID, item.ProductID, item.Quantity, item.Price) // Assuming Price is part of CartItem
		if err != nil {
			stmt.Close()
			return err
		}
		stmt.Close()
	}

	return tx.Commit()
}

func (os *OrderStoreImpl) GetOrders(ctx context.Context, userID int) ([]models.OrderDetail, error) {
	stmt, err := os.db.PrepareContext(ctx, `
		SELECT OrderID, TotalAmount, Status, CreatedAt
		FROM Orders
		WHERE UserID = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.OrderDetail
	for rows.Next() {
		var order models.OrderDetail
		if err := rows.Scan(&order.ID, &order.Total, &order.Status, &order.CreatedAt); err != nil {
			return nil, err
		}

		items, err := os.getOrderItems(ctx, order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}

func (os *OrderStoreImpl) GetOrder(ctx context.Context, orderID int) (*models.OrderDetail, error) {
	stmt, err := os.db.PrepareContext(ctx, `
		SELECT UserID, TotalAmount, Status, CreatedAt
		FROM Orders
		WHERE OrderID = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	order := models.OrderDetail{ID: orderID}
	if err := stmt.QueryRowContext(ctx, orderID).Scan(&order.UserID, &order.Total, &order.Status, &order.CreatedAt); err != nil {
		return nil, err
	}

	items, err := os.getOrderItems(ctx, orderID)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return &order, nil
}

func (os *OrderStoreImpl) getOrderItems(ctx context.Context, orderID int) ([]models.OrderItem, error) {
	stmt, err := os.db.PrepareContext(ctx, `
		SELECT ProductID, Quantity, Price
		FROM OrderItems
		WHERE OrderID = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (os *OrderStoreImpl) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	stmt, err := os.db.PrepareContext(ctx, `
		UPDATE Orders
		SET Status = $1
		WHERE OrderID = $2
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, status, orderID)
	return err
}
