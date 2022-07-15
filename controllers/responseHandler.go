package controllers

import (
	"accountanApp/models"
	"encoding/json"
	"net/http"
)

func sendUnauthorizedResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	var response models.Response
	response.Message = "Unauthorized Access"
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var response models.Response
	response.Message = message
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response models.Response
	response.Message = "success"
	response.Data = data
	json.NewEncoder(w).Encode(response)
}
