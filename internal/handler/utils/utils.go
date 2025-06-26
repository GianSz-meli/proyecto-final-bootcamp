package utils

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
}

func SendError(w http.ResponseWriter, apiError models.ApiError) {
	//errResponse := models.ErrorResponse{
	//	Message:    apiError.Message,
	//	StatusCode: apiError.StatusCode,
	//}
	SendJsonResponse(w, apiError.StatusCode, apiError)
}
