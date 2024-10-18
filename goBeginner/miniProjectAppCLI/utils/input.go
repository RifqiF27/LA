package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func InputStr(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", errors.New("input bukan string yang valid")
	}
	return input, nil
}
func InputInt(prompt string) (int, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("input bukan angka")
	}
	return value, nil
}

func InputFloat(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("input bukan angka desimal")
	}
	return value, nil
}
func InputBool(prompt string) (bool, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return false, errors.New("gagal membaca input")
	}

	value, err := strconv.ParseBool(input)

	if err != nil {
		if input == "available" {
			return true, nil
		} else if input == "unavailable" {
			return false, nil
		}
		return false, errors.New("input harus 'true', 'false', 'available', atau 'unavailable'")
	}

	return value, nil
}
