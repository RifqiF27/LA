package utils

import (
	"encoding/json"
	"net/http"
	"todo/model"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := model.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
