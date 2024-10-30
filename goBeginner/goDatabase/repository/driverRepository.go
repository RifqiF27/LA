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

func (r *OrderRepoDb) GetAllOrderDriver(orders *[]model.OrderSummary) error {
	query := `
	SELECT
    d."name" AS "driver_name",
    TO_CHAR ("date_order", 'YYYY-MM-DD') AS "month",
    COUNT(o."id") AS "total_order"
	FROM
    "Order" o
    JOIN "Driver" d ON o."driver_id" = d."id"
	GROUP BY
    "month",
    "driver_name"
	ORDER BY
    "month",
    "total_order" DESC`

	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var summary model.OrderSummary

		if err := rows.Scan(&summary.DriverName, &summary.Month, &summary.TotalOrder); err != nil {
			return err
		}

		*orders = append(*orders, summary)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
