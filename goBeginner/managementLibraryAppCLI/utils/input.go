package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func InputStr(prompt string) (string, error) {
	green := "\033[32m"
	reset := "\033[0m"

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s%s%s", green, prompt, reset)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("\033[31mInput tidak valid\033[0m")
	}
	input = strings.TrimSpace(input)

	return input, nil
}
