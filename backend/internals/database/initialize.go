package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func init_tables(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Initialization started...")
	defer fmt.Println("Initialization finished.")

	// initialize migration
	_, err := db.ExecContext(ctx, migration_init)
	if err != nil {
		return err
	}

	migrationID := ""
	row := db.QueryRowContext(ctx, "Select VersionNumber from MigrationVersion LIMIT 1;")
	err = row.Scan(&migrationID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if migrationID != "01" {
		fmt.Println("Migrating Version 01")
		_, err := db.ExecContext(ctx, migration_1)
		if err != nil {
			fmt.Println("Error migrating One : ", err)
			return err
		}
	}

	return nil
}

const (
	migration_init = `
	-- Create migration version table
	CREATE TABLE IF NOT EXISTS MigrationVersion (
		VersionID SERIAL PRIMARY KEY,
		VersionNumber VARCHAR(50) UNIQUE NOT NULL,
		MigrationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		Description TEXT
	);
	`

	migration_1 = `
	-- Down Migration: Drop tables and indexes
	DROP TABLE IF EXISTS CartItems;
	DROP TABLE IF EXISTS Cart;
	DROP TABLE IF EXISTS OrderItems;
	DROP TABLE IF EXISTS Orders;
	DROP TABLE IF EXISTS Products;
	DROP TABLE IF EXISTS Users;

	DROP INDEX IF EXISTS idx_products_name;
	DROP INDEX IF EXISTS idx_orders_userid;
	DROP INDEX IF EXISTS idx_unique_email;
	-- DROP TABLE IF EXISTS MigrationVersion;

	-- Up Migration: Create tables and indexes
	CREATE TABLE Users (
		UserID SERIAL PRIMARY KEY,
		Email VARCHAR(255) UNIQUE NOT NULL,
		Password VARCHAR(255) NOT NULL,
		Role VARCHAR(50) CHECK (Role IN ('Customer', 'Admin'))
	);

	CREATE UNIQUE INDEX idx_unique_email ON Users (Email);

	CREATE TABLE Products (
		ProductID SERIAL PRIMARY KEY,
		Name VARCHAR(255) NOT NULL,
		Description TEXT,
		Price DECIMAL(10, 2) NOT NULL,
		Stock INT NOT NULL,
		ImageURL VARCHAR(255)
	);

	CREATE INDEX idx_products_name ON Products (Name);

	CREATE TABLE Orders (
		OrderID SERIAL PRIMARY KEY,
		UserID INT NOT NULL REFERENCES Users(UserID) ON DELETE CASCADE,
		TotalAmount DECIMAL(10, 2) NOT NULL,
		Status VARCHAR(50) NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX idx_orders_userid ON Orders (UserID);

	CREATE TABLE OrderItems (
		OrderItemID SERIAL PRIMARY KEY,
		OrderID INT NOT NULL REFERENCES Orders(OrderID) ON DELETE CASCADE,
		ProductID INT NOT NULL REFERENCES Products(ProductID) ON DELETE CASCADE,
		Quantity INT NOT NULL,
		Price DECIMAL(10, 2) NOT NULL
	);

	CREATE TABLE Cart (
		CartID SERIAL PRIMARY KEY,
		UserID INT NOT NULL REFERENCES Users(UserID) ON DELETE CASCADE
	);

	CREATE TABLE CartItems (
		CartItemID SERIAL PRIMARY KEY,
		CartID INT NOT NULL REFERENCES Cart(CartID) ON DELETE CASCADE,
		ProductID INT NOT NULL REFERENCES Products(ProductID) ON DELETE CASCADE,
		Quantity INT NOT NULL
	);

	-- Optionally, insert the migration version record after successful execution
	INSERT INTO MigrationVersion (VersionNumber, Description)
	VALUES ('01', 'Initial migration with Users, Products, Orders, OrderItems, Cart, CartItems tables and unique indexes.')
	ON CONFLICT (VersionNumber) DO NOTHING;  -- To avoid inserting duplicate version records
	`
)
