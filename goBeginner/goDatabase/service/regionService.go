package service

import (
	"database/sql"
	"errors"
	"fmt"
	"main/model"
	"main/repository"
)


func InputDataRegion(db *sql.DB, name string) error {

	if name == "" {
		return errors.New("data region tidak boleh kosong")
	}

	
	// tx, err = db.Begin()
	// if err != nil {
	// 	return fmt.Errorf("gagal memulai transaksi: %w", err)
	// }

	regionRepo := repository.NewRegionRepo(db)
	region := model.Region{
		Name: name,
	}

	err := regionRepo.Create(&region)
	if err != nil {
		// tx.Rollback()
		return fmt.Errorf("gagal membuat region: %w", err)
	}

	fmt.Println("berhasil input data region dengan id ", region.ID)
	return nil
}

func GetAllRegion(db *sql.DB) error  {
	regionRepo := repository.NewRegionRepo(db)
	var regions []model.Region

	err := regionRepo.GetAll(&regions)
	if err != nil {
		return fmt.Errorf("gagal mengambil data region: %w", err)
	}
	for _, region := range regions {
		fmt.Printf("ID: %d, name: %s\n", region.ID, region.Name)
		
	}
	return nil
}
func GetAllRegionOrderSummary(db *sql.DB) error  {
	regionRepo := repository.NewRegionRepo(db)
	var regions []model.RegionSummary

	err := regionRepo.GetAllOrderRegion(&regions)
	if err != nil {
		return fmt.Errorf("gagal mengambil data region: %w", err)
	}
	for _, region := range regions {
		fmt.Printf("Name: %s, Area: %s, Month: %s, Total Order: %d\n", region.Name, region.Area, region.Month, region.TotalOrder)
		
	}
	return nil
}
