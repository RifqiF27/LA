package service

import (
	"ecommerce/model"
	"ecommerce/repository"
	"fmt"
	"math"

	"go.uber.org/zap"
)

type ShippingServiceInterface interface {
	Create(customer *model.Shipping) error
	GetAll() (*[]model.Shipping, error)
	ShippingCost(payload model.RequestDestination) (*model.CostResponse, error)
	CreateNewShipping(orderID string, shippingID uint, originLatLong string, destinationLatLong string, totalPayment float64) (model.OrderShipping, error)
	TrackDelivery(orderShippingID uint) (model.HistoryDelivery, error)
	UpdateShippingStatus(orderShippingID uint, status string, location string) (model.HistoryDelivery, error)
}

type ShippingService struct {
	Repo repository.AllRepository
	Log  *zap.Logger
}

func NewShippingService(repo repository.AllRepository, log *zap.Logger) ShippingServiceInterface {
	return &ShippingService{
		Repo: repo,
		Log:  log,
	}
}

func (shippingService *ShippingService) Create(customer *model.Shipping) error {
	return nil
}

func (shippingService *ShippingService) GetAll() (*[]model.Shipping, error) {
	shippings, err := shippingService.Repo.ShippingRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all shipping costs: %w", err)
	}
	return &shippings, nil
}

func (service *ShippingService) ShippingCost(payload model.RequestDestination) (*model.CostResponse, error) {
	shippingData, err := service.Repo.ShippingRepo.GetByID(payload.ShippingID)
	if err != nil {
		return nil, fmt.Errorf("failed to get shipping data: %w", err)
	}

	distance, err := service.Repo.ShippingRepo.GetDestination(payload)
	if err != nil {
		return nil, err
	}

	cost := shippingData.Price
	if distance != nil {
		if *distance > 1500000 {
			cost += 10
		}
	}

	if payload.Qty > 2 {
		cost += 5
	}
	costResponse := &model.CostResponse{
		Shipping: shippingData.Name,
		Distance: math.Round(*distance / 1000),
		Cost:     cost,
	}

	return costResponse, nil
}

func (service *ShippingService) CreateNewShipping(orderID string, shippingID uint, originLatLong string, destinationLatLong string, totalPayment float64) (model.OrderShipping, error) {
	return service.Repo.ShippingRepo.CreateNewShipping(orderID, shippingID, originLatLong, destinationLatLong, totalPayment)
}

func (service *ShippingService) TrackDelivery(orderShippingID uint) (model.HistoryDelivery, error) {
	return service.Repo.ShippingRepo.TrackDelivery(orderShippingID)
}

func (service *ShippingService) UpdateShippingStatus(orderShippingID uint, status string, location string) (model.HistoryDelivery, error) {
	return service.Repo.ShippingRepo.UpdateShippingStatus(orderShippingID, status, location)
}
