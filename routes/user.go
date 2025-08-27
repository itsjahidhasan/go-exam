package routes

import (
	"go-exam/handler"
	"net/http"
)

func UserRoutes(r *http.ServeMux) {
	// Placeholder for user routes
	r.HandleFunc(("GET /users"), handler.UserGetAll)
}
