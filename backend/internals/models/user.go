package models

import "time"

type AuthToken struct {
	Exp    time.Time `json:"exp"`
	UserID int       `json:"user_id"`
	Role   string    `json:"role"`
}

func (a *AuthToken) Validate() bool {
	return time.Now().Before(a.Exp)
}

type UserDetails struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
	Gender   string `json:"gender"`
	DOB      string `json:"birth_date"`
}

type LoginRequestPayload struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
