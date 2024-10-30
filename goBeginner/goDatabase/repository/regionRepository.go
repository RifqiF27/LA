package repository

import (
	"database/sql"
	"main/model"
)

type RegionRepository interface {
	Create(region *model.Region) error
	GetAll(region *[]model.Region) error
	GetAllOrderRegion(region *[]model.RegionSummary) error
	
}
type RegionRepoDb struct {
	DB *sql.DB
}

func NewRegionRepo(db *sql.DB) RegionRepository {
	return &RegionRepoDb{DB: db}
}

func (r *RegionRepoDb) Create(region *model.Region) error {
	query := `INSERT INTO "Region" (name) VALUES ($1) RETURNING id`
	err := r.DB.QueryRow(query, region.Name).Scan(&region.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RegionRepoDb) GetAll(regions *[]model.Region) error {
	query := `SELECT id, customer_id, driver_id, date_Region, status FROM "Region"`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var region model.Region
		if err := rows.Scan(&region.ID, &region.Name); err != nil {
			return err
		}
		*regions = append(*regions, region)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
func (r *RegionRepoDb) GetAllOrderRegion(regions *[]model.RegionSummary) error {
	query := `
	SELECT
    r."name" AS "region_name",
    ra."area" AS "region_area",
    TO_CHAR (o."date_order", 'YYYY-MM') AS "month",
    COUNT(DISTINCT o."id") AS "total_order"
	FROM
    "Order" o
    JOIN "Region_Area" ra ON o."region_area_id" = ra."id"
    JOIN "Region" r ON ra."region_id" = r."id"
	where
    "status" = true
	GROUP BY
    r."name",
    ra."area",
    "month"
	ORDER by
    "month" desc,
    "total_order" DESC
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var region model.RegionSummary
		if err := rows.Scan(&region.Name, &region.Area, &region.Month, &region.TotalOrder); err != nil {
			return err
		}
		*regions = append(*regions, region)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}