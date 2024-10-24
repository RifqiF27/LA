package main

import (
	"context"
	"fmt"
	"main/challenge/controllers"
	"main/challenge/models"
	// "time"
)

func main() {
	for {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		
		// fmt.Println(models.Items.Session, ">>>>>>>>")
		loginSuccess := controllers.Login(ctx)
		
		if loginSuccess {
			
			for models.Items.Session {
				controllers.MainMenu()
				// time.Sleep(500 * time.Millisecond)
			}

			fmt.Println("Session expired. Returning to login page...")
		} else {

			fmt.Println("Login failed. Please try again.")
		}
	}
}
