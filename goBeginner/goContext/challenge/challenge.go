package main

import (
	"context"
	"main/challenge/controllers"
	"main/challenge/models"
)

func main() {
	for {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if models.Items.Session {
			select {
			case <-ctx.Done():
				break
			default:
				controllers.MainMenu()
			}
		} else {
			controllers.Login(ctx)
		}
		
	}
}
