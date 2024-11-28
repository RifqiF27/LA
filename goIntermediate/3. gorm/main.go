package main

import (
	"ecommerce/database"
	"ecommerce/infra"
	"ecommerce/router"
	"log"

	"go.uber.org/zap"
)

func main() {
	ctx, err := infra.NewContext()
	if err != nil {
		ctx.Log.Panic("Error", zap.Error(err))
		return
	}

	log.Println("Starting migration...")
	database.MigrateTables(ctx.DB)
	log.Println("Migration completed successfully.")

	log.Println("Starting seeding...")
	database.SeedData(ctx.DB)
	log.Println("Seeding completed successfully.")

	router.SetupReouter(ctx)
}
