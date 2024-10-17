package model

import (
	"fmt"
)

type Account struct {
	Name    string
	Email   string
	Saldo Saldo 
}

func NewAcc(name, email string) Account {
	if name == "" || email == "" {
		panic("Error: Nama dan Email tidak boleh kosong!")
	}
	return Account{Name: name, Email: email, Saldo: NewSaldo()} 
}

func AddAcc(accounts []Account, a ...Account) []Account {
	return append(accounts, a...) 
}

func PrintAccounts(accounts []Account) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	fmt.Println("Daftar Akun:")
	for i, account := range accounts {
		fmt.Printf("%d. Akun berhasil ditambahkan {nama: %s, email: %s, saldo: %d}\n", i+1, account.Name, account.Email, account.Saldo.Saldo)
	}
}
