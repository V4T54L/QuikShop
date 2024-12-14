package handler

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, statusCode int, obj any) {
	w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(obj)
	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}

func errorResponse(w http.ResponseWriter, statusCode int, err string) {
	jsonResponse(w, statusCode, map[string]string{
		"error": err,
	})
}

func messageResponse(w http.ResponseWriter, statusCode int, msg string) {
	jsonResponse(w, statusCode, map[string]string{
		"message": msg,
	})
}
