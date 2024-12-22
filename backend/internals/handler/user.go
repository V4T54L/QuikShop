package handler

import (
	"backend/internals/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.LoginRequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.userService.LoginUser(r.Context(), payload)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, models.TokenResponse{Token: token})
}

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.UserDetails
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.userService.SignupUser(r.Context(), &payload)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusCreated, user)
}
