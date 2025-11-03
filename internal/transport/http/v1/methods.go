package v1

import (
	"debez/internal/models"
	"debez/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	users, err := h.service.GetUsers(r.Context(), offset, limit)
	if err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SaveUser(w http.ResponseWriter, r *http.Request) {
	var userC models.UserCreate

	if err := json.NewDecoder(r.Body).Decode(&userC); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	user, err := h.service.SaveUser(r.Context(), &userC)
	if err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	var userID models.UserID

	if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetUser(r.Context(), &userID)
	if err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userU models.UserUpdate

	if err := json.NewDecoder(r.Body).Decode(&userU); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	if err := h.service.UpdateUser(r.Context(), &userU); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userID models.UserID

	if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	if err := h.service.DeleteUser(r.Context(), &userID); err != nil {
		utils.SendError(w, &models.ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
}
