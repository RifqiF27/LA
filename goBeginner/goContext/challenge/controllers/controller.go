package controllers

import (
	"bufio"
	"context"
	"fmt"
	"main/challenge/models"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func Login(ctx context.Context) {

	fmt.Println("Please login:")

	time.Sleep(1 * time.Second)
	models.Items.Session = true

	sessionCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	fmt.Println("Login successful! You have 20 seconds to complete your actions.")

	<-sessionCtx.Done()
	models.Items.Session = false
	fmt.Println("Session expired. Please login again.")
	os.Exit(0)

}

func MainMenu() {
	fmt.Println("\n--- Main Menu ---")
	fmt.Println("1. Show Products")
	fmt.Println("2. View Cart")
	fmt.Println("3. Checkout")
	fmt.Println("4. Exit")

	choice := getUserInput("Choose an option: ")
	ClearScreen()
	switch choice {
	case "1":
		showProducts()
	case "2":
		viewCart()
	case "3":
		checkout()
	case "4":
		models.Items.Session = false
	default:
		fmt.Println("Invalid option. Try again.")
	}
}

func showProducts() {
	fmt.Println("\n--- Available Products ---")
	for _, product := range models.Products {
		fmt.Printf("%d. %s - IDR %.2f\n", product.ID, product.Name, product.Price)
	}

	productID := getUserInput("Enter the product ID to add to cart (or type 'back' to return): ")
	if strings.ToLower(productID) == "back" {
		return
	}
	quantity := getUserInput("Enter the quantity: ")

	for _, product := range models.Products {
		if fmt.Sprintf("%d", product.ID) == productID {
			qty := 1
			fmt.Sscanf(quantity, "%d", &qty)
			models.Items.Cart = append(models.Items.Cart, models.CartItem{Product: product, Quantity: qty})
			fmt.Println("Product added to cart.")
			return
		}
	}

	fmt.Println("Product not found.")
}

func viewCart() {
	fmt.Println("\n--- Your Cart ---")
	if len(models.Items.Cart) == 0 {
		fmt.Println("Your cart is empty.")
		return
	}

	total := 0.0
	for _, item := range models.Items.Cart {
		fmt.Printf("%s x%d - IDR %.2f\n", item.Product.Name, item.Quantity, item.Product.Price*float64(item.Quantity))
		total += item.Product.Price * float64(item.Quantity)
	}
	fmt.Printf("Total: $%.2f\n", total)
}

func checkout() {
	if len(models.Items.Cart) == 0 {
		fmt.Println("Your cart is empty. Add some products before checking out.")
		return
	}

	viewCart()
	confirm := getUserInput("Proceed to checkout? (yes/no): ")

	if strings.ToLower(confirm) == "yes" {
		models.Items.Cart = []models.CartItem{}
		fmt.Println("Checkout successful! Your items will be shipped soon.")
	} else {
		fmt.Println("Checkout cancelled.")
	}
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
