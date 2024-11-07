package utils

import (
	"book-store/collections"
	"encoding/json"
	"fmt"
)

func SendJSONResponse(status string, message string, data interface{}) {
	response := collections.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
