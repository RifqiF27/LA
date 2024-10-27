package utils

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInput        = errors.New("input invalid")
	ErrInputLenGreaterThan = errors.New("input must be greater than or equal 3")
	ErrInputMustBeLetters  = errors.New("input must be Letters")
	ErrInputMustBeNumber   = errors.New("input must be Number")
	ErrInputGreaterThan    = errors.New("input must be greater than or equal to 5000")
	ErrEditAccount         = errors.New("edit account failed")
	ErrUsernameExists      = errors.New("username already exist")
	ErrInvalidLogin        = errors.New("username or password invalid")
	ErrInsufficientFunds   = errors.New("insufficient balance")
	ErrAccountNotFound     = errors.New("account not found")
	ErrNotLoggedIn         = errors.New("please login first")
	ErrSelfTransfer        = errors.New("can't transfer to own account")
)

func ErrHandler(err error) {

	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidInput):
			fmt.Println(Red + "Error: Input invalid." + Reset)

		case errors.Is(err, ErrInputLenGreaterThan):
			fmt.Println(Red + "Error: Input must be greater than or equal 3." + Reset)

		case errors.Is(err, ErrInputMustBeNumber):
			fmt.Println(Red + "Error: Input must be number." + Reset)

		case errors.Is(err, ErrInputMustBeLetters):
			fmt.Println(Red + "Error: Input must be letters." + Reset)

		case errors.Is(err, ErrInputGreaterThan):
			fmt.Println(Red + "Error: Input must be greater than or equal to 5000." + Reset)

		case errors.Is(err, ErrEditAccount):
			fmt.Println(Red + "Error: Edit account failed." + Reset)

		case errors.Is(err, ErrUsernameExists):
			fmt.Println(Red + "Error: Username already exists, choose another username." + Reset)

		case errors.Is(err, ErrInvalidLogin):
			fmt.Println(Red + "Error: Username or password invalid." + Reset)

		case errors.Is(err, ErrInsufficientFunds):
			fmt.Println(Red + "Error: Your balance is insufficient to make this transaction." + Reset)

		case errors.Is(err, ErrAccountNotFound):
			fmt.Println(Red + "Error: Account not found, make sure the username is correct." + Reset)

		case errors.Is(err, ErrNotLoggedIn):
			fmt.Println(Red + "Error: Please login first." + Reset)

		case errors.Is(err, ErrSelfTransfer):
			fmt.Println(Red + "Error: Can't transfer to own account." + Reset)

		default:
			fmt.Println(Red + "Error: An error occurred, please try again." + Reset)
		}
	}
}

