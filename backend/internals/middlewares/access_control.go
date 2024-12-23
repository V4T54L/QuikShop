package middlewares

import (
	"backend/internals/utils"
	"context"
	"net/http"
)

type UserKey string

const (
	UserIDKey UserKey = "userID"
)

func UserOnlyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check for token in request header
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) <= len("Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr = tokenStr[len("Bearer "):]
		payload, err := utils.VerifyJWT(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, payload.ID)
		r = r.WithContext(ctx)

		// call next handler
		h.ServeHTTP(w, r)
	})
}

func AdminOnlyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check for token in request header
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) <= len("Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr = tokenStr[len("Bearer "):]
		payload, err := utils.VerifyJWT(tokenStr)
		if err != nil || payload.Role != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// call next handler
		h.ServeHTTP(w, r)
	})
}
