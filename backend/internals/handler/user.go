package handler

import (
	"backend/internals/models"
	"backend/internals/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user models.UserDetail
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Register the user
	if err := h.service.RegisterUser(r.Context(), &user); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusCreated, user)
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserDetail
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.LoginUser(r.Context(), user.Email, user.Password)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	user, err := h.service.GetUserProfile(r.Context(), userID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, user)
}
