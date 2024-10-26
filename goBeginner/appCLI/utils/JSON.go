package utils

import (
    "encoding/json"
    "os"
)

func ReadFile(fileName string, data interface{}) error {
    file, err := os.ReadFile(fileName)
    if err != nil {
        // ErrHandler(err)
        return err
    }
    return json.Unmarshal(file, data)
}

func WriteFile(fileName string, data interface{}) error {
    file, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        // ErrHandler(err)
        return err
    }
    return os.WriteFile(fileName, file, 0644)
}
