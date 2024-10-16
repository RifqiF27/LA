package main

import "fmt"

// buatkan sebuah struct(nama, score,) buatkan fungsi untuk merubah score

type Value struct {
	nama  string
	score int
}

func (v *Value) changeScore(newScore int) {
	v.score = newScore

}

// buatkan sebuah struct lalu lakukan inisialisasikan dan buatkan 3 data dari struct tsb
type Person struct {
	nama        string
	address     string
	phoneNumber string
}

func main() {
	name1 := Value{nama: "lumo", score: 1}

	fmt.Println("before", name1)

	name1.changeScore(5)
	// name1 = Value{nama: "lumo", score: 2}

	fmt.Println("after", name1)

	// buatkan sebuah struct lalu lakukan inisialisasikan dan buatkan 3 data dari struct tsb

	var data []Person
	data = append(data, Person{nama: "lomoshive", address: "jakarta", phoneNumber: "021775213"})
	data = append(data, Person{nama: "lumo", address: "jakarta", phoneNumber: "021775213"})
	data = append(data, Person{nama: "academy", address: "jakarta", phoneNumber: "021775213"})

	fmt.Println(data)
	for _, p := range data {
		fmt.Println("Name:", p.nama)
		fmt.Println("Address:", p.address)
		fmt.Println("Phone Number:", p.phoneNumber)
	}
}
