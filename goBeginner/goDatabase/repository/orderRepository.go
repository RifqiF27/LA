package repository

import (
	"database/sql"
	"main/model"
)

type OrderRepository interface {
	Create(order *model.Order) error
	GetAll(order *[]model.Order) error
	GetAllOrderTotal(order *[]model.OrderSummary) error
	GetAllOrderTotalCustomer(order *[]model.OrderSummary) error
}
type OrderRepoDb struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) OrderRepository {
	return &OrderRepoDb{DB: db}
}

func (r *OrderRepoDb) Create(order *model.Order) error {
	query := `INSERT INTO "Order" (customer_id, driver_id,  region_area_id, date_order, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(query, order.CustomerId, order.DriverId,order.RegionAreaId, order.DateOrder, order.Status).Scan(&order.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepoDb) GetAll(orders *[]model.Order) error {
	query := `SELECT id, customer_id, driver_id, region_area_id, date_order, status FROM "Order"`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.CustomerId, &order.DriverId,&order.RegionAreaId, &order.DateOrder, &order.Status); err != nil {
			return err
		}
		*orders = append(*orders, order)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepoDb) GetAllOrderTotal(orders *[]model.OrderSummary) error {
	query := `
	SELECT
    TO_CHAR ("date_order", 'YYYY-MM') AS "month",
    COUNT("id") AS "total_order"
	FROM
    "Order"
	WHERE
    "status" = TRUE
	GROUP BY
    "month"
	order by
    "month" desc`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var summary model.OrderSummary

		if err := rows.Scan(&summary.Month, &summary.TotalOrder); err != nil {
			return err
		}

		*orders = append(*orders, summary)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepoDb) GetAllOrderTotalCustomer(orders *[]model.OrderSummary) error {
	query := `
	SELECT
	c."name" AS "customer_name",
	TO_CHAR("date_order", 'YYYY-MM') AS "month",
	COUNT(o."id") AS "total_order"
	FROM
	"Order" o
	JOIN "Customer" c ON o."customer_id" = c."id"
	WHERE
	"status" = TRUE
	GROUP BY
	"month",
	"customer_name"
	ORDER BY
	"month",
	"total_order" DESC;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var summary model.OrderSummary

		if err := rows.Scan(&summary.CustomerName, &summary.Month, &summary.TotalOrder); err != nil {
			return err
		}

		*orders = append(*orders, summary)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
func (r *OrderRepoDb) GetAllOrderTime(orders *[]model.OrderSummary) error {
	query := `
	SELECT
    TO_CHAR ("dateOrder", 'HH24:MI:SS') AS "hour",
    COUNT("id") AS "total_order"
	FROM
    "Order"
	GROUP BY
    "hour"
	ORDER BY
    "total_order" DESC`

	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var summary model.OrderSummary

		if err := rows.Scan(&summary.Month, &summary.TotalOrder); err != nil {
			return err
		}

		*orders = append(*orders, summary)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
