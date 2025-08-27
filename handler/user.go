package handler

import (
	"go-exam/repository"
	"go-exam/utils"
	"net/http"
)

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	userRepo := &repository.UserRepo{}

	users, err := userRepo.GetAll(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	utils.WriteJSON(w, http.StatusOK, users)
}
