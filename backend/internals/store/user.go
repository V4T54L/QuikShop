package store

import (
	"backend/internals/models"
	"context"
)

type UserStore interface {
	AddUser(ctx context.Context, user *models.UserDetails) (*models.UserDetails, error)
	GetUserByUsername(ctx context.Context, username string) (*models.UserDetails, error)
}
