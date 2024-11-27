package repository

import (
	"database/sql"
	"ecommerce/helper"
	"ecommerce/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type ShippingRepoInterface interface {
	Create(customer *model.Shipping) error
	GetAll() (*[]model.Shipping, error)
	GetDestination(destination model.RequestDestination) (*float64, error)

	GetByID(id int) (*model.Shipping, error)
}

type ShippingRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewShippingRepository(db *sql.DB, log *zap.Logger) ShippingRepoInterface {
	return &ShippingRepository{
		DB:     db,
		Logger: log,
	}
}

func (shippingRepo *ShippingRepository) GetAll() (*[]model.Shipping, error) {
	var shippings []model.Shipping
	query := "SELECT id, name, cost FROM shipping_costs"

	rows, err := shippingRepo.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch shipping costs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var shipping model.Shipping
		if err := rows.Scan(&shipping.ID, &shipping.Name, &shipping.Price); err != nil {
			return nil, fmt.Errorf("failed to scan shipping row: %w", err)
		}
		shippings = append(shippings, shipping)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	return &shippings, nil
}

func (shippingRepo *ShippingRepository) GetByID(id int) (*model.Shipping, error) {
	var shipping model.Shipping
	query := "SELECT id, name, cost FROM shipping_costs WHERE id = $1"

	row := shippingRepo.DB.QueryRow(query, id)
	if err := row.Scan(&shipping.ID, &shipping.Name, &shipping.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no shipping found with id %d", id)
		}
		return nil, fmt.Errorf("failed to fetch shipping by ID: %w", err)
	}

	return &shipping, nil
}

func (shippingRepo *ShippingRepository) Create(shipping *model.Shipping) error {
	query := `INSERT INTO shipping_costs (name, cost) VALUES (?, ?)`
	_, err := shippingRepo.DB.Exec(query, shipping.Name, shipping.Price)
	if err != nil {
		return fmt.Errorf("failed to insert new shipping cost: %w", err)
	}
	return nil
}

func (shippingRepo *ShippingRepository) GetDestination(destination model.RequestDestination) (*float64, error) {

	originLat := 39.8283
	originLon := -98.5795

	query := `
        SELECT 
            address->>'latitude' AS destination_lat, 
            address->>'longitude' AS destination_lon
        FROM 
            users 
        WHERE id = $1;
    `

	var destinationLatStr, destinationLonStr string
	row := shippingRepo.DB.QueryRow(query, destination.UserID)

	if err := row.Scan(&destinationLatStr, &destinationLonStr); err != nil {
		shippingRepo.Logger.Error("Error fetching coordinates from users table", zap.Error(err))
		return nil, errors.New("failed to fetch coordinates from the users table")
	}

	destinationLat, err := strconv.ParseFloat(destinationLatStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse destination latitude: %w", err)
	}

	destinationLon, err := strconv.ParseFloat(destinationLonStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse destination longitude: %w", err)
	}

	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=false",
		originLon, originLat, destinationLon, destinationLat)

	var header http.Header
	data, err := helper.HTTPRequest("GET", header, url, nil)
	if err != nil {
		shippingRepo.Logger.Error("Error making HTTP request to OSRM API", zap.Error(err))
		return nil, errors.New("failed to fetch route data from OSRM API")
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		shippingRepo.Logger.Error("Error unmarshalling response from OSRM API", zap.Error(err))
		return nil, errors.New("failed to decode API response")
	}

	routes, ok := dataMap["routes"].([]interface{})
	if !ok || len(routes) == 0 {
		shippingRepo.Logger.Error("OSRM API response does not contain valid routes", zap.Any("response", dataMap))
		return nil, errors.New("no valid routes found in the OSRM response")
	}

	route, ok := routes[0].(map[string]interface{})
	if !ok {
		shippingRepo.Logger.Error("OSRM API response route is not in expected format", zap.Any("route", routes[0]))
		return nil, errors.New("invalid route format in OSRM response")
	}

	distance, ok := route["distance"].(float64)
	if !ok {
		shippingRepo.Logger.Error("OSRM API response does not contain valid distance", zap.Any("route", route))
		return nil, errors.New("distance not found in OSRM route")
	}

	return &distance, nil
}
