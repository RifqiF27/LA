package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadSession() (map[string]interface{}, error) {
	sessionData, err := os.ReadFile("session.json")
	if err != nil {
		return nil, fmt.Errorf("gagal membaca session: %w", err)
	}

	var session map[string]interface{}
	err = json.Unmarshal(sessionData, &session)
	if err != nil {
		return nil, fmt.Errorf("gagal mendekode session: %w", err)
	}

	return session, nil
}