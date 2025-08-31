package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, code int, message string) {
	WriteJSON(w, code, map[string]string{"error": message})
}

func ResponseModifier(d any, m string, s bool) map[string]any {
	if !s {
		return map[string]any{
			"error":   d,
			"message": m,
			"status":  s,
		}
	}
	return map[string]any{
		"data":    d,
		"message": m,
		"status":  s,
	}
}
