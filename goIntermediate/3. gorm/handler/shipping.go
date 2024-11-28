package handler

import (
	"ecommerce/helper"
	"ecommerce/model"
	"ecommerce/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShippingHandler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewShippingHandler(service service.AllService, log *zap.Logger) ShippingHandler {
	return ShippingHandler{
		Service: service,
		Log:     log,
	}
}

func (h *ShippingHandler) Create(ca *gin.Context) {

}

func (h *ShippingHandler) GetAllShipping(c *gin.Context) {

	shippings, err := h.Service.CustomerService.GetAll()
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	result := gin.H{
		"shipping_costs": shippings,
	}
	helper.SuccessResponseWithData(c, "All shipping costs fetched successfully", http.StatusOK, result)
}

func (h *ShippingHandler) ShippingCost(c *gin.Context) {
	var data model.RequestDestination
	if err := c.ShouldBindQuery(&data); err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	cost, err := h.Service.CustomerService.ShippingCost(data)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
	}

	result := gin.H{
		"cost": cost,
	}

	helper.SuccessResponseWithData(c, "success", http.StatusOK, result)
}

func (h *ShippingHandler) CreateNewShippingHandler(c *gin.Context) {
	var req struct {
		OrderID            string  `json:"order_id"`
		ShippingID         uint    `json:"shipping_id"`
		OriginLatLong      string  `json:"origin_latlong"`
		DestinationLatLong string  `json:"destination_latlong"`
		TotalPayment       float64 `json:"total_payment_shipping"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	orderShipping, err := h.Service.CustomerService.CreateNewShipping(req.OrderID, req.ShippingID, req.OriginLatLong, req.DestinationLatLong, req.TotalPayment)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	result := gin.H{
		"order_shipping": orderShipping,
	}

	helper.SuccessResponseWithData(c, "success", http.StatusCreated, result)

}

func (h *ShippingHandler) TrackDeliveryHandler(c *gin.Context) {
	orderShippingID, err := strconv.ParseUint(c.Param("order_shipping_id"), 10, 32)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	history, err := h.Service.CustomerService.TrackDelivery(uint(orderShippingID))
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	result := gin.H{
		"status":   history.Status,
		"location": history.Location,
	}

	helper.SuccessResponseWithData(c, "success", http.StatusOK, result)

}

func (h *ShippingHandler) UpdateShippingStatusHandler(c *gin.Context) {
	var req struct {
		Status   string `json:"status"`
		Location string `json:"location"`
	}

	orderShippingID, err := strconv.ParseUint(c.Param("order_shipping_id"), 10, 32)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	history, err := h.Service.CustomerService.UpdateShippingStatus(uint(orderShippingID), req.Status, req.Location)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	result := gin.H{
		"history": history,
	}

	helper.SuccessResponseWithData(c, "success", http.StatusOK, result)
}
