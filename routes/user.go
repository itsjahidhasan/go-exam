package routes

import (
	"database/sql"
	"go-exam/handler"
	"go-exam/repository"
	"net/http"
)

func UserRoutes(r *http.ServeMux, c *sql.DB) {
	// Initialize repository and inject it into handlers
	userRepo := repository.NewUserRepo(c)
	handler.SetUserRepo(userRepo)
	// Placeholder for user routes
	r.HandleFunc(("GET /users"), handler.UserGetAll)
}
