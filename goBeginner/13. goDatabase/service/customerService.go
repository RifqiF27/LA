package service

import (
	"database/sql"
	"errors"
	"fmt"
	"main/model"
	"main/repository"
)

func InputDataCustomer(tx *sql.Tx, name, phoneNumber, address string, userID int) error {
	if name == "" {
		return errors.New("name tidak boleh kosong")
	}
	if phoneNumber == "" {
		return errors.New("phoneNumber tidak boleh kosong")
	}
	if address == "" {
		return errors.New("address tidak boleh kosong")
	}

	// tx, err := db.Begin()
	// if err != nil {
	// 	return fmt.Errorf("gagal memulai transaksi: %w", err)
	// }

	custRepo := repository.NewCustomerRepo(tx)
	customer := model.Customer{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		UserID:      userID,
	}

	err := custRepo.Create(&customer)
	if err != nil {
		// tx.Rollback()
		return fmt.Errorf("gagal membuat user: %w", err)
	}

	fmt.Println("berhasil input data customer dengan id ", customer.ID)
	return nil
}
