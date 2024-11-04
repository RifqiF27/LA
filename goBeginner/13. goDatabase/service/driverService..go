package service

import (
	"database/sql"
	"errors"
	"fmt"
	"main/model"
	"main/repository"
)

func InputDataDriver(tx *sql.Tx, name, phoneNumber, address, vehicle string, userID int) error {
	if name == "" {
		return errors.New("name tidak boleh kosong")
	}
	if phoneNumber == "" {
		return errors.New("phoneNumber tidak boleh kosong")
	}
	if address == "" {
		return errors.New("address tidak boleh kosong")
	}
	if vehicle == "" {
		return errors.New("vehicle tidak boleh kosong")
	}

	// tx, err := db.Begin()
	// if err != nil {
	// 	return fmt.Errorf("gagal memulai transaksi: %w", err)
	// }

	driverRepo := repository.NewDriverRepo(tx)
	driver := model.Driver{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		Vehicle:     vehicle,
		UserID:      userID,
	}

	err := driverRepo.Create(&driver)
	if err != nil {
		// tx.Rollback() 
		return fmt.Errorf("gagal membuat user: %w", err)
	}

	fmt.Println("berhasil input data driver dengan id ", driver.ID)
	return nil
}

func GetAllSummaryOrdersDriver(db *sql.DB) error {
	orderRepo := repository.NewOrderRepo(db)
	var orders []model.OrderSummary

	err := orderRepo.GetAllOrderDriver(&orders)
	if err != nil {
		return fmt.Errorf("gagal mengambil data: %w", err)
	}

	for _, order := range orders {
		fmt.Printf("Driver Name: %s, Month: %s, Total Order: %d\n",order.DriverName, order.Month, order.TotalOrder)
	}
	return nil
}