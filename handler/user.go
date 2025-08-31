package handler

import (
	"encoding/json"
	"go-exam/modals"
	"go-exam/repository"
	"go-exam/utils"
	"io"
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

func UserGetAllPurchaseHistoryWithRestaurant(w http.ResponseWriter, r *http.Request) {
	id := utils.GetPathParams(w, r.URL.Path)
	users, err := userRepo.GetUserPurchaseHistory(r.Context(), id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	utils.WriteJSON(w, http.StatusOK, users)
}
func PurchaseDish(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ResponseModifier(modals.UserPurchaseResponse{}, "Unable to read request body", false))
	}
	defer r.Body.Close()

	var requestBody modals.UserPurchaseRequest
	if err := json.Unmarshal(body, &requestBody); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ResponseModifier(modals.UserPurchaseResponse{}, "Invalid json body", false))

	}

	res, rqErr := userRepo.PurchaseDish(r.Context(), requestBody)
	if rqErr != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]any{"error": rqErr.Error()})
	}

	utils.WriteJSON(w, http.StatusCreated, utils.ResponseModifier(res, "Data successfully fetched", true))
}
