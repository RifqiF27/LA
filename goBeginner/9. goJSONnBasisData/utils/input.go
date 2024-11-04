package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func InputInt() (int, error) {
	var input string
	// fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.Atoi(input)
	if err != nil || value <= 0 {
		return 0, errors.New("input must be number")
	}
	return value, nil
}
