package service

import (
	"database/sql"
	"errors"
	"fmt"
	"main/model"
	"main/repository"
)

func ValidateUserIdExists(db *sql.DB, table string, id int) (bool, error) {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM "%s" WHERE id=$1)`, table)
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("gagal memvalidasi %s ID: %w", table, err)
	}
	return exists, nil
}

func InputDataOrder(db *sql.DB, customerId, driverId, regionAreaId int, dateOrder string, status bool) error {

	// Validasi customerId
	exists, err := ValidateUserIdExists(db, "Customer", customerId)
	if err != nil {
		return fmt.Errorf("gagal memvalidasi customer ID: %w", err)
	}
	if !exists {
		return errors.New("customer ID tidak valid atau tidak ditemukan")
	}

	// Validasi driverId
	exists, err = ValidateUserIdExists(db, "Driver", driverId)
	if err != nil {
		return fmt.Errorf("gagal memvalidasi driver ID: %w", err)
	}
	if !exists {
		return errors.New("driver ID tidak valid atau tidak ditemukan")
	}

	exists, err = ValidateUserIdExists(db, "Driver", regionAreaId)
	if err != nil {
		return fmt.Errorf("gagal memvalidasi region area ID: %w", err)
	}
	if !exists {
		return errors.New("region area ID tidak valid atau tidak ditemukan")
	}

	if dateOrder == "" {
		return errors.New("date order tidak boleh kosong")
	}

	// tx, err = db.Begin()
	// if err != nil {
	// 	return fmt.Errorf("gagal memulai transaksi: %w", err)
	// }

	orderRepo := repository.NewOrderRepo(db)
	order := model.Order{
		CustomerId:   customerId,
		DriverId:     driverId,
		RegionAreaId: regionAreaId,
		DateOrder:    dateOrder,
		Status:       status,
	}

	err = orderRepo.Create(&order)
	if err != nil {
		// tx.Rollback()
		return fmt.Errorf("gagal membuat user: %w", err)
	}

	fmt.Println("berhasil input data order dengan id ", order.ID)
	return nil
}

func GetAllOrders(db *sql.DB) error {
	orderRepo := repository.NewOrderRepo(db)
	var orders []model.Order

	err := orderRepo.GetAll(&orders)
	if err != nil {
		return fmt.Errorf("gagal mengambil data order: %w", err)
	}
	for _, order := range orders {
		fmt.Printf("ID: %d, customer_id: %d, driver_id: %d,region area: %d, date_order: %s, status: %v\n", order.ID, order.CustomerId, order.DriverId, order.RegionAreaId, order.DateOrder, order.Status)

	}
	return nil
}

func GetAllSummaryOrdersTotal(db *sql.DB) error {
	orderRepo := repository.NewOrderRepo(db)
	var orders []model.OrderSummary

	err := orderRepo.GetAllOrderTotal(&orders)
	if err != nil {
		return fmt.Errorf("gagal mengambil data order: %w", err)
	}

	for _, order := range orders {
		fmt.Printf("Month: %s, Total Order: %d\n", order.Month, order.TotalOrder)
	}
	return nil
}
func GetAllSummaryOrdersCustomer(db *sql.DB) error {
	orderRepo := repository.NewOrderRepo(db)
	var orders []model.OrderSummary

	err := orderRepo.GetAllOrderTotalCustomer(&orders)
	if err != nil {
		return fmt.Errorf("gagal mengambil data order: %w", err)
	}

	for _, order := range orders {
		fmt.Printf("Customer Name: %s, Month: %s, Total Order: %d\n", order.CustomerName, order.Month, order.TotalOrder)
	}
	return nil
}
func GetAllSummaryOrdersTime(db *sql.DB) error {
	orderRepo := repository.NewOrderRepo(db)
	var orders []model.OrderSummary

	err := orderRepo.GetAllOrderTotalCustomer(&orders)
	if err != nil {
		return fmt.Errorf("gagal mengambil data order: %w", err)
	}

	for _, order := range orders {
		fmt.Printf("Hours: %s, Total Order: %d\n", order.Month, order.TotalOrder)
	}
	return nil
}
