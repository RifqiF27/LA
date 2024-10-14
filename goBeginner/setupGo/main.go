package main

import (
	"fmt"
	// "strconv"
)

func main() {
	// s := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// avg := float32(s[2]+s[5]+s[9]) / float32(4.0)
	// fmt.Println(avg)

	// a:= 30
	// b:= float32(24.5)
	// c:=	-45
	// d:= float32(0.67)

	// result := float32(a)+b*float32(c)/d
	// var f2 float64 = float64(result)
	// var hasil string = strconv.FormatFloat(f2, 'f', 6, 64)
	// fmt.Println(hasil)

	hargaJual := 150000
	hargaBeli := 100000
	operational := 1000
	disc := 15
	terjual := 100

	hargaDisc := float64(hargaJual - (hargaJual*disc)/100)
	fmt.Println(hargaDisc)

	total := hargaDisc * float64(terjual)
	fmt.Printf("%.f\n", total)

	totalBiaya := (hargaBeli + operational) * terjual
	fmt.Println(totalBiaya)

	totalKeuntungan := total - float64(totalBiaya)
	fmt.Printf("%.f\n", totalKeuntungan)
}
