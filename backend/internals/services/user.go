package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"backend/utils"
	"context"
)

type UserService struct {
	store store.UserStore
}

func NewUserService(store store.UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) LoginUser(mainCtx context.Context, creds models.LoginRequestPayload) (string, error) {
	ctx, cancel := context.WithTimeout(mainCtx, maxQueryTime)
	defer cancel()

	user, err := s.store.GetUserByUsername(ctx, creds.Username)
	if err != nil {
		return "", err
	}
	tokenStr, err := utils.GetTokenString(user.ID, user.Role)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s *UserService) SignupUser(mainCtx context.Context, user *models.UserDetails) (*models.UserDetails, error) {
	ctx, cancel := context.WithTimeout(mainCtx, maxQueryTime)
	defer cancel()

	return s.store.AddUser(ctx, user)
}
