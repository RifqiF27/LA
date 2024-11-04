package model

import (
	"errors"
	"fmt"
)

type Saldo struct {
	Saldo int 
    
}

func NewSaldo() Saldo {
	return Saldo{Saldo: 0} 
}
func HandleErrSaldo(val int) error{
    if val <= 0 {
        return errors.New("Jumlah yang ditambahkan harus lebih dari 0")
    }
    return nil
}

func (s *Saldo) Debit(saldo int, accounts []Account) {
	
    errHandle := HandleErrSaldo(saldo)
    if errHandle != nil {
        fmt.Println(errHandle)
        return
    }
	s.Saldo += saldo
	fmt.Printf("Saldo berhasil ditambahkan %d ", s.Saldo)
    PrintAccounts(accounts)

}


func (s *Saldo) Credit(saldo int, accounts []Account) {
    errHandle := HandleErrSaldo(saldo)
    if errHandle != nil {
        fmt.Println(errHandle)
        return
    }
	if s.Saldo < saldo {
		panic("Saldo tidak cukup.")
	}
    
	s.Saldo -= saldo
	fmt.Printf("Saldo berhasil dikurangi %d ", s.Saldo)
    PrintAccounts(accounts)
}


// func (s Saldo) PrintSaldo() {
// 	fmt.Printf("Saldo saat ini adalah: %d\n", s.Saldo)
// }
