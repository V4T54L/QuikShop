package store

import (
	"backend/internals/models"
	"context"
	"database/sql"
	"fmt"
)

type UserStore interface {
	Create(ctx context.Context, user *models.UserDetail) error
	GetByUserID(ctx context.Context, userID int) (*models.UserDetail, error)
	GetByEmail(ctx context.Context, email string) (*models.UserDetail, error)
}

type userStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *userStore {
	return &userStore{db}
}

func (s *userStore) Create(ctx context.Context, user *models.UserDetail) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.PrepareContext(ctx, `
    INSERT INTO Users (Email, Password, Role) 
    VALUES ($1, $2, $3) RETURNING UserID
    `)

	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, user.Email, user.Password, user.Role).Scan(&user.ID); err != nil {
		return fmt.Errorf("could not execute statement: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	user.Password = ""

	fmt.Printf("User created with ID: %d\n", user.ID)
	return nil
}

func (s *userStore) GetByEmail(ctx context.Context, email string) (*models.UserDetail, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT UserID, Role from Users 
	WHERE Email=$1`)
	if err != nil {
		return nil, fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	user := models.UserDetail{Email: email}
	if err := stmt.QueryRowContext(ctx, user.Email).Scan(&user.ID, &user.Role); err != nil {
		return nil, fmt.Errorf("could not execute statement: %w", err)
	}

	return &user, nil
}

func (s *userStore) GetByUserID(ctx context.Context, userID int) (*models.UserDetail, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT Email, Role from Users 
	WHERE UserID=$1`)
	if err != nil {
		return nil, fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	user := models.UserDetail{ID: userID}
	if err := stmt.QueryRowContext(ctx, user.ID).Scan(&user.Email, &user.Role); err != nil {
		return nil, fmt.Errorf("could not execute statement: %w", err)
	}

	return &user, nil
}
