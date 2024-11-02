package repository

import (
	"database/sql"
	"service-inventory/model"
)

type LocationRepository interface {
	Create(location *model.Location) (*model.Location, error)
	
}

type LocationRepoDb struct {
	DB *sql.DB
}

func NewLocationRepo(db *sql.DB) LocationRepository {
	return &LocationRepoDb{DB: db}
}


func (r *LocationRepoDb) Create(location *model.Location) (*model.Location, error) {
	query := `INSERT INTO "Locations" (location_name) VALUES ($1) RETURNING location_id`
	err := r.DB.QueryRow(query, location.LocationName).Scan(&location.ID)
	if err != nil {
		return nil, err
	}

	return location, nil
}


