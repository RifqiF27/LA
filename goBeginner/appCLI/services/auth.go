package services

import (
	"fmt"
	"main/models"
	"main/utils"
	"regexp"
	"strings"
)

func Login(userService *models.UserServices, accountService *models.AccountService) {

	for {
		fmt.Println("== Login ==")
		username, err := utils.InputStr("Username: ")
		if err != nil {
			utils.ErrHandler(err)
			continue
		}

		password, err := utils.InputStr("Password: ")
		if err != nil {
			utils.ErrHandler(err)
			continue
		}

		err = userService.Login(username, password)
		if err != nil {
			utils.ErrHandler(utils.ErrInvalidLogin)
			choice, err := utils.InputStr("Create your e-wallet account (yes/any): ")
			if err == nil && strings.ToLower(choice) == "yes" {
				register(userService)
			} else {
				continue
			}
		} else {
			accountService.CreateAccount(username)
			fmt.Println(utils.Green + "Login successfully!" + utils.Reset)
			break
		}
	}
}

func register(userService *models.UserServices) {
	for {

		fmt.Println("== Register ==")
		name, err := utils.InputStr("Name: ")
		if err != nil {
			utils.ErrHandler(err)
			continue
		} else {
			isString := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

			if !isString(name) {
				utils.ErrHandler(utils.ErrInputMustBeLetters)
				continue
			}
		}
		phoneNumber, err := utils.InputStr("Phone Number: ")
		if err != nil {
			utils.ErrHandler(err)
			continue
		} else {
			isString := regexp.MustCompile(`^[0-9]+$`).MatchString

			if !isString(phoneNumber) {
				utils.ErrHandler(utils.ErrInputMustBeNumber)
				continue
			}
		}
		username, err := utils.InputStr("Username: ")
		if err != nil {
			utils.ErrHandler(err)
			continue
		}
		password, err := utils.InputStr("Password: ")
		if err != nil {
			utils.ErrHandler(err)
			continue
		}

		e := userService.Register(name, phoneNumber, username, password)
		if e != nil {
			utils.ErrHandler(e)
		} else {
			fmt.Println(utils.Green + "Registration successfully, Please login!" + utils.Reset)
			break
		}
	}
}
