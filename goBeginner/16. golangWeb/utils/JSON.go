package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func DecodeJSONFile(filename string, out interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(out)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	return nil
}

func WriteJSONFile(filename string, data interface{}) error {
	sessionJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("gagal membuat data sesi: %w", err)
	}

	err = os.WriteFile(filename, sessionJSON, 0644)
	if err != nil {
		return fmt.Errorf("gagal menyimpan sesi: %w", err)
	}

	return nil
}

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