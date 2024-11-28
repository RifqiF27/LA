package database

import (
	"ecommerce/model"
	"log"

	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) {
	log.Println("Starting migration...")

	err := db.AutoMigrate(
		&model.Shipping{},
		&model.OrderShipping{},
		&model.HistoryDelivery{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}
