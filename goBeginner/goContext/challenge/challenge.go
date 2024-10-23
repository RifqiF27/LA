package main

import (
	"context"
	"main/challenge/controllers"
	"main/challenge/models"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go controllers.Login(ctx)

	time.Sleep(1 * time.Second)
	for models.Items.Session {
		controllers.MainMenu()
	}
}
