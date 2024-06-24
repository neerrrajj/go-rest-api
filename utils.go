package main

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// DOUBT : can't we just return if there is not token instead of an empty string?

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")
	if tokenAuth != "" {
		return tokenAuth
	}
	if tokenQuery != "" {
		return tokenQuery
	}
	return ""
}