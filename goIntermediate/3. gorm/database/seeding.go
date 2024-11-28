package database

import (
	"ecommerce/model"
	"log"

	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {

	tx := db.Begin()

	var count int64
	if err := tx.Model(&model.Shipping{}).Count(&count).Error; err != nil {
		tx.Rollback()
		log.Fatalf("Error checking shipping data: %v", err)
		return
	}

	if count > 0 {
		tx.Rollback()
		log.Println("Seeding skipped, data already exists.")
		return
	}

	shippings := []model.Shipping{
		{Name: "Standard Shipping", Price: 5.00},
		{Name: "Express Shipping", Price: 15.00},
		{Name: "Overnight Shipping", Price: 25.00},
	}

	if err := tx.Create(&shippings).Error; err != nil {
		tx.Rollback()
		log.Fatalf("Seeding failed: %v", err)
	}

	tx.Commit()
	log.Println("Seeding completed successfully with sample data.")
}
