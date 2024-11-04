package utils

import (
	"fmt"
	// "regexp"
	"strconv"
)

func InputStr(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", ErrInvalidInput
	} else if len(input) < 3 {
		return "", ErrInputLenGreaterThan

	}

	// isString := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	// if !isString(input) {
	// 	return "", errors.New("input must be letters")
	// }
	return input, nil
}
func InputInt(prompt string) (int, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, ErrInvalidInput
	} 
	return value, nil
}

func InputFloat(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, ErrInvalidInput
	} else if value < 5000 {
		return 0, ErrInputGreaterThan
	}
	return value, nil
}
