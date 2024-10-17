package model

import "fmt"

type Saldo struct {
	Saldo int 
    
}

func NewSaldo() Saldo {
	return Saldo{Saldo: 0} 
}


func (s *Saldo) Debit(saldo int) {
	if saldo <= 0 {
		panic("Error: Jumlah yang ditambahkan harus lebih dari 0.")
	}
	s.Saldo += saldo
	fmt.Printf("Saldo berhasil ditambahkan, saldo sekarang: %d\n", s.Saldo)
}


func (s *Saldo) Credit(saldo int) {
	if saldo <= 0 {
		panic("Error: Jumlah yang dikurangi harus lebih dari 0.")
	}
	if s.Saldo < saldo {
		panic("Error: Saldo tidak cukup untuk pengurangan ini.")
	}
	s.Saldo -= saldo
	fmt.Printf("Saldo berhasil dikurangi, saldo sekarang: %d\n", s.Saldo)
}


func (s Saldo) PrintSaldo() {
	fmt.Printf("Saldo saat ini adalah: %d\n", s.Saldo)
}
