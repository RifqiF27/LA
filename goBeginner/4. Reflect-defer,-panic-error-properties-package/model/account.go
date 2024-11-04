package model

import (
	"errors"
	"fmt"
)

type Account struct {
	Name  string
	Email string
	Saldo Saldo
}


func NewAcc(name, email string) Account {
	
	return Account{Name: name, Email: email, Saldo: NewSaldo()}
}
func HandleErr(name, email string) error{
	if name == "" || email == ""{
		return errors.New("Nama dan Email tidak boleh kosong")
	}
	return nil
}

func AddAcc(accounts []Account, a ...Account) []Account {
	
	handleErr := HandleErr(a[0].Name, a[0].Email)
	if len(a) == 1 {
		if handleErr != nil {
			fmt.Println(handleErr)
			return accounts
		}
	}
	accounts = append(accounts, a...)
	if len(a) > 1 {
		return accounts
	}
	
	fmt.Printf("Akun berhasil ditambahkan ")
	PrintAccounts(accounts)
	return accounts
}

func PrintAccounts(accounts []Account) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from error:", r)
	// 	}
	// }()

	fmt.Printf("%+v\n", accounts)

	// fmt.Println("Daftar Akun:")
	// for i, account := range accounts {
	// 	fmt.Printf("%d. Akun berhasil ditambahkan {nama: %s, email: %s, saldo: %d}\n", i+1, account.Name, account.Email, account.Saldo.Saldo)
	// }
}
