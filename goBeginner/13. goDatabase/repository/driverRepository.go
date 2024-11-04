package repository

import (
	"database/sql"
	"main/model"
)

type DriverRepository interface {
	Create(driver *model.Driver) error
}

type DriverRepoDb struct {
	TX *sql.Tx
}

func NewDriverRepo(tx *sql.Tx) DriverRepository {
	return &DriverRepoDb{TX: tx}
}

func (r *DriverRepoDb) Create(driver *model.Driver) error {
	query := `INSERT INTO "Driver" (name, phone_number, address, vehicle, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.TX.QueryRow(query, driver.Name, driver.PhoneNumber, driver.Address, driver.Vehicle, driver.UserID).Scan(&driver.ID)
	if err != nil {
		return err
	}
	return nil
}

