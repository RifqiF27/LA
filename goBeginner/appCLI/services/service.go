package services

import (
	"fmt"
	"main/models"
	"main/utils"
	"regexp"
	"strings"
)

func TopUp(accountService *models.AccountService, userService *models.UserServices) {
	amount, err := utils.InputFloat("Amount Top-Up: ")
	if err != nil {
		utils.ErrHandler(err)
		return
	}

	description, err := utils.InputStr("Description Top-Up: ")
	if err != nil {
		utils.ErrHandler(err)
		return
	}
	user := userService.GetLoggedInUser()
	accountService.TopUp(user.Username, amount, description)
	fmt.Println(utils.Green + "Top-Up berhasil!" + utils.Reset)
}

func Transfer(accountService *models.AccountService, userService *models.UserServices) {
	toUsername, err := utils.InputStr("Transfer to (Username): ")
	if err != nil {
		utils.ErrHandler(err)
		return
	}
	amount, err := utils.InputFloat("Amount Transfer: ")
	if err != nil {
		utils.ErrHandler(err)
		return
	}

	fromUser := userService.GetLoggedInUser()
	e := accountService.Transfer(fromUser.Username, toUsername, amount)
	if e != nil {
		utils.ErrHandler(e)
	} else {
		fmt.Println(utils.Green + "Transfer berhasil!" + utils.Reset)
	}
}

func EditAccount(userService *models.UserServices) {

	newName, err := utils.InputStr("New Name: ")
	if err != nil {
		utils.ErrHandler(err)
		return
	} else {
		isString := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if !isString(newName) {
			utils.ErrHandler(utils.ErrInputMustBeLetters)
			return
		}
	}

	newPhoneNumber, err := utils.InputStr("New Phone Number: ")
	if err != nil {
		utils.ErrHandler(err)
		return
	} else {
		isString := regexp.MustCompile(`^[0-9]+$`).MatchString

		if !isString(newPhoneNumber) {
			utils.ErrHandler(utils.ErrInputMustBeNumber)
			return
		}
	}
	password, err := utils.InputStr("New Password: ")
	if err != nil {
		utils.ErrHandler(err)
		return
	}

	e := userService.EditAccount(newName, newPhoneNumber, password)
	if e != nil {
		utils.ErrHandler(e)
	} else {
		fmt.Println(utils.Green + "Account has been updated" + utils.Reset)
	}
}

func CheckBalance(accountService *models.AccountService, userService *models.UserServices) {
	user := userService.GetLoggedInUser()
	balance, _ := accountService.GetBalance(user.Username)
	fmt.Printf("Nama: %s\nSaldo Anda: \033[32m\033[1m%.2f\033[0m\n", utils.Bold+user.Name+utils.Reset, balance)
}

func HistoryTransaction(accountService *models.AccountService, userService *models.UserServices) {
	user := userService.GetLoggedInUser()
	transactions := accountService.GetTransactions(user.Username)

	fmt.Println("History Transaction:")
	fmt.Printf(utils.BrightPurple + "%-5s %-15s %-15s %-15s %-10s %-15s\n" + utils.Reset,"ID", "From", "To", "Amount", "Type", "Desc")
	fmt.Println(strings.Repeat("-", 80))
	if len(transactions) == 0 {
		fmt.Println("no transactions yet")
	} else {
		for _, t := range transactions {
			var color string
			if t.Type == "transfer" && t.FromUserID == user.Username {
				color = utils.Red
			} else {
				color = utils.Green
			}

			fmt.Printf("%s%-5d %-15s %-15s %-15.2f %-10s %-15s\033[0m\n",
				color, t.ID, t.FromUserID, t.ToUserID, t.Amount, t.Type, t.Description)
			fmt.Println(strings.Repeat("-", 80))

		}
	}
}
