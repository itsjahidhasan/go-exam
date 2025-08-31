package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func GetPathParams(w http.ResponseWriter, p string) int64 {
	parts := strings.Split(p, "/")
	if len(parts) < 3 {
		WriteJSON(w, http.StatusBadRequest, ResponseModifier("Invalid routes", "", false))
	}
	id, err := strconv.ParseInt(parts[2], 10, 0)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ResponseModifier(err.Error(), "", false))
	}
	return id
}
