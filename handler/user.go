package handler

import (
	"go-exam/repository"
	"go-exam/utils"
	"net/http"
)

var userRepo *repository.UserRepo

// Use this function to inject the repo
func SetUserRepo(repo *repository.UserRepo) {
	userRepo = repo
}

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := userRepo.GetAll(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	utils.WriteJSON(w, http.StatusOK, users)
}
