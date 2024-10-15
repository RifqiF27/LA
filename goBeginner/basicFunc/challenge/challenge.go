package main

import (
	"fmt"
)

func priceProduct(merk string) int {
	switch merk {
	case "Adidas":
		price := 200000
		return price
	case "Kappa":
		price := 600000
		return price
	case "Puma":
		price := 150000
		return price
	default:
		price := 0
		return price
	}
}

func priceDisc(product string, product2 string) int {
	if product == "Adidas" && product2 == "Puma" {
		return (priceProduct("Adidas") + priceProduct("Puma") - 50000)
	} else if product == "Puma" && product2 == "Kappa" {
		return (priceProduct("Puma") + priceProduct("Kappa") - 150000)
	} else if product == "Adidas" && product2 == "Kappa" {
		return (priceProduct("Adidas") + priceProduct("Kappa") - 75000)
	} else {
		price := priceProduct(product)
		if price > 0 {
			return price
		}
		return priceProduct(product2)

	}
}

func main() {
	product := "Adidas"
	product2 := "Puma"
	fmt.Println("Result --> Rp.", priceDisc(product, product2))
}
