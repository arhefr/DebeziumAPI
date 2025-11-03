package utils

import (
	"debez/internal/models"
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, mes *models.ErrorMessage, statusCode int) error {

	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(mes)
}
