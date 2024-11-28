package repository

import (
	"ecommerce/helper"
	"ecommerce/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ShippingRepoInterface interface {
	Create(customer *model.Shipping) error
	GetAll() ([]model.Shipping, error)
	GetDestination(destination model.RequestDestination) (*float64, error)
	GetByID(id int) (*model.Shipping, error)
	CreateNewShipping(orderID string, shippingID uint, originLatLong string, destinationLatLong string, totalPayment float64) (model.OrderShipping, error)
	TrackDelivery(orderShippingID uint) (model.HistoryDelivery, error)
	UpdateShippingStatus(orderShippingID uint, status string, location string) (model.HistoryDelivery, error)
}

type ShippingRepository struct {
	DB     *gorm.DB
	DB2    *gorm.DB
	Logger *zap.Logger
}

func NewShippingRepository(db *gorm.DB, db2 *gorm.DB, log *zap.Logger) ShippingRepoInterface {
	return &ShippingRepository{
		DB:     db,
		DB2:    db2,
		Logger: log,
	}
}

func (shippingRepo *ShippingRepository) GetAll() ([]model.Shipping, error) {
	var shippings []model.Shipping
	if err := shippingRepo.DB.Find(&shippings).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch shipping costs: %w", err)
	}
	return shippings, nil
}

func (shippingRepo *ShippingRepository) GetByID(id int) (*model.Shipping, error) {
	var shipping model.Shipping
	if err := shippingRepo.DB.First(&shipping, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no shipping found with id %d", id)
		}
		return nil, fmt.Errorf("failed to fetch shipping by ID: %w", err)
	}
	return &shipping, nil
}

func (shippingRepo *ShippingRepository) Create(shipping *model.Shipping) error {
	if err := shippingRepo.DB.Create(shipping).Error; err != nil {
		return fmt.Errorf("failed to insert new shipping cost: %w", err)
	}
	return nil
}

func (shippingRepo *ShippingRepository) GetDestination(destination model.RequestDestination) (*float64, error) {
	originLat := 39.8283
	originLon := -98.5795

	var user struct {
		Latitude  string `gorm:"column:latitude"`
		Longitude string `gorm:"column:longitude"`
	}
	err := shippingRepo.DB2.Raw(`
		SELECT address->>'latitude' AS latitude, 
			   address->>'longitude' AS longitude 
		FROM users WHERE id = ?`, destination.UserID).Scan(&user).Error
	if err != nil {
		shippingRepo.Logger.Error("Error fetching coordinates from users table", zap.Error(err))
		return nil, errors.New("failed to fetch coordinates from the users table")
	}

	destinationLat, err := strconv.ParseFloat(user.Latitude, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse destination latitude: %w", err)
	}

	destinationLon, err := strconv.ParseFloat(user.Longitude, 64)
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

func (shippingRepo *ShippingRepository) CreateNewShipping(orderID string, shippingID uint, originLatLong string, destinationLatLong string, totalPayment float64) (model.OrderShipping, error) {

	orderShipping := model.OrderShipping{
		OrderID:            orderID,
		ShippingID:         shippingID,
		OriginLatLong:      originLatLong,
		DestinationLatLong: destinationLatLong,
		TotalPayment:       totalPayment,
	}

	if err := shippingRepo.DB.Create(&orderShipping).Error; err != nil {
		return model.OrderShipping{}, fmt.Errorf("failed to create new shipping: %w", err)
	}

	if err := shippingRepo.DB.Preload("Shipping").First(&orderShipping, orderShipping.ID).Error; err != nil {
		return model.OrderShipping{}, fmt.Errorf("failed to preload shipping: %w", err)
	}

	history := model.HistoryDelivery{
		OrderShippingID: orderShipping.ID,
		Status:          "pending",
		Location:        "Origin Location",
	}

	if err := shippingRepo.DB.Create(&history).Error; err != nil {
		return model.OrderShipping{}, fmt.Errorf("failed to create initial delivery history: %w", err)
	}

	return orderShipping, nil
}

func (shippingRepo *ShippingRepository) TrackDelivery(orderShippingID uint) (model.HistoryDelivery, error) {
	var history model.HistoryDelivery
	if err := shippingRepo.DB.Where("order_shipping_id = ?", orderShippingID).Order("created_at desc").First(&history).Error; err != nil {
		return model.HistoryDelivery{}, fmt.Errorf("failed to track delivery: %w", err)
	}
	return history, nil
}

func (shippingRepo *ShippingRepository) UpdateShippingStatus(orderShippingID uint, status string, location string) (model.HistoryDelivery, error) {

	history := model.HistoryDelivery{
		OrderShippingID: orderShippingID,
		Status:          status,
		Location:        location,
	}

	if err := shippingRepo.DB.Create(&history).Error; err != nil {
		return model.HistoryDelivery{}, fmt.Errorf("failed to update shipping status: %w", err)
	}
	if err := shippingRepo.DB.Preload("OrderShipping.Shipping").First(&history, history.ID).Error; err != nil {
		return model.HistoryDelivery{}, fmt.Errorf("failed to preload shipping: %w", err)
	}

	return history, nil
}
