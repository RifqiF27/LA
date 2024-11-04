package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func InputStr(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", errors.New("input tidak valid")
	}
	isString := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if !isString(input) {
		return "", errors.New("input harus berupa huruf")
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
	} else if value <= 0 {
		return 0, errors.New("input harus lebih dari 0")
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
	} else if value <= 0 {
		return 0, errors.New("input harus lebih dari 0")
	}
	return value, nil
}

// func InputBool(prompt string) (bool, error) {
// 	var input string
// 	fmt.Print(prompt)
// 	_, err := fmt.Scanln(&input)
// 	if err != nil {
// 		return false, errors.New("gagal membaca input")
// 	}

// 	value, err := strconv.ParseBool(input)

// 	if err != nil {
// 		if input == "available" {
// 			return true, nil
// 		} else if input == "unavailable" {
// 			return false, nil
// 		}
// 		return false, errors.New("input harus 'true', 'false', 'available', atau 'unavailable'")
// 	}

// 	return value, nil
// }
