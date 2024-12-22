package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"backend/internals/utils"
	"context"
	"errors"
)

type UserService struct {
	store store.UserStore
}

func NewUserService(store store.UserStore) *UserService {
	return &UserService{store}
}

func (s *UserService) RegisterUser(ctx context.Context, user *models.UserDetail) error {
	// Validate user details
	if user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("email and password are required")
	}

	// Hash the password
	user.Password = utils.HashPassword(user.Password)
	return s.store.Create(ctx, user)
}

func (s *UserService) LoginUser(ctx context.Context, email, password string) (string, error) {
	// Validate user details
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := s.store.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// Compare the password
	if !utils.ComparePassword(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	payload := models.TokenPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	token, err := utils.GenerateJWT(&payload)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, userID int) (*models.UserDetail, error) {
	return s.store.GetByUserID(ctx, userID)
}
