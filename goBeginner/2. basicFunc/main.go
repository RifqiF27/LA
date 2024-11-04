package main

import "fmt"

// buatkan code golang untuk mengkonversi dari angka(1-7) jadi nama hari
func main()  {
	angka := "7"
	switch angka{
	case "1":
		fmt.Println("senin")
	case "2":
		fmt.Println("selasa")
	case "3":
		fmt.Println("rabu")
	case "4":
		fmt.Println("kamis")
	case "5":
		fmt.Println("jumat")
	case "6":
		fmt.Println("sabtu")
	case "7":
		fmt.Println("minggu")
	default:
		fmt.Println("not found")
	}

	var nilai int = 100
	if nilai < 0 || nilai > 100 {
		fmt.Println("Error") 
	} else if nilai >= 80 && nilai <= 100 {
        fmt.Println("A")
    } else if nilai >= 70 {
        fmt.Println("B")
    } else if nilai >= 60 {
        fmt.Println("C")
    } else if nilai >= 50 {
        fmt.Println("D")
    } else if nilai >= 0 && nilai < 50 {
        fmt.Println("E")
    } 
}