package main

import (
	"fmt"
	"main/model"
)

var accounts []model.Account

func init() {
	fmt.Println("Sistem keuangan dimulai...")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic caught in main:", r)
		}
	}()

	acc1 := model.NewAcc("Lumo", "lumo@mail.com")
	acc2 := model.NewAcc("Aca", "aca@mail.com")

	accounts = model.AddAcc(accounts, acc1, acc2)

	model.PrintAccounts(accounts)

	accounts[0].Saldo.Debit(1000)
	accounts[0].Saldo.Credit(500)

	accounts[1].Saldo.Debit(5000)
	accounts[1].Saldo.Credit(1500)

	accounts[0].Saldo.PrintSaldo()
	accounts[1].Saldo.PrintSaldo()

	fmt.Printf("%+v\n", accounts)
}
