package store

import (
	"backend/internals/models"
	"context"
	"errors"
)

type MockUserStore struct {
	users     []models.UserDetails
	usernames map[string]int
}

func NewMockUserStore() *MockUserStore {
	users := []models.UserDetails{
		{
			ID: 1, Name: "Admin User", Username: "admin", Password: "Admin@123",
			Role: "admin", Gender: "Male", DOB: "01/01/1993",
		},
		{
			ID: 2, Name: "Customer User", Username: "customer", Password: "123",
			Role: "user", Gender: "Male", DOB: "01/11/1997",
		},
	}
	return &MockUserStore{users: users, usernames: make(map[string]int)}
}

func (s *MockUserStore) GetUserByUsername(ctx context.Context, username string) (*models.UserDetails, error) {
	userIdx, ok := s.usernames[username]
	if !ok {
		return nil, errors.New("user does not exist")
	}
	user := s.users[userIdx]
	user.Password = ""
	return &user, nil
}

func (s *MockUserStore) AddUser(ctx context.Context, user *models.UserDetails) (*models.UserDetails, error) {
	if _, ok := s.usernames[user.Username]; ok {
		return nil, errors.New("username is taken, please choose another one")
	}
	idx := len(s.users)
	m_user := models.UserDetails{ID: idx + 1, Name: user.Name, Password: user.Password, Username: user.Username, Role: "user", Gender: user.Gender, DOB: user.DOB}
	s.usernames[m_user.Username] = idx
	s.users = append(s.users, m_user)
	updated_user := m_user
	updated_user.Password = ""
	return &updated_user, nil
}
