package views

import (
    "fmt"
    "main/models"
    "main/services"
    "main/utils"
)

func MainMenu() {
    userService := models.NewUser("users.json", "session.json")
    accountService := models.NewAccountService("accounts.json", "transactions.json")

    
    for {
        if !userService.IsLoggedIn() {
            services.Login(userService, accountService)
        }
        fmt.Println(utils.Bold+"\n==== E-Wallet ===="+utils.Reset)
        fmt.Println("1. Check Balance")
        fmt.Println("2. Top-Up")
        fmt.Println("3. Transfer")
        fmt.Println("4. History Transaction")
        fmt.Println("5. Edit Account")
        fmt.Println("6. Logout")
        fmt.Println("7. Exit")

        choice,_ := utils.InputInt("Select Menu: ")
		
        utils.ClearScreen()
        switch choice {
        case 1:
            services.CheckBalance(accountService, userService)
        case 2:
            services.TopUp(accountService, userService)
        case 3:
            services.Transfer(accountService, userService)
        case 4:
            services.HistoryTransaction(accountService, userService)
        case 5:
            services.EditAccount(userService)
        case 6:
            userService.Logout()
            fmt.Println(utils.Green+"Logout successfully!"+utils.Reset)
            services.Login(userService, accountService)
        case 7:
            fmt.Println(utils.Bold+"Exiting App"+utils.Reset)
            return
        default:
            fmt.Println(utils.Yellow+"Input invalid, please try again!"+utils.Reset)
        }
    }
}
