package utils

import (
	"encoding/json"
	"fmt"
	"main/model"
	"os"
)

func SendJSONResponse(statusCode int, message string, data interface{}) {
	response := model.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = os.WriteFile("body.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan body:", err)
		return
	}

	fmt.Println(string(jsonData))
}
