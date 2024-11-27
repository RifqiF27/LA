package handler

import (
	"ecommerce/helper"
	"ecommerce/model"
	"ecommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShippingHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewShippingHandler(service service.AllService, log *zap.Logger) ShippingHadler {
	return ShippingHadler{
		Service: service,
		Log:     log,
	}
}

func (shippingHadler *ShippingHadler) Create(ca *gin.Context) {

}

func (shippingHandler *ShippingHadler) GetAllShipping(c *gin.Context) {

	shippings, err := shippingHandler.Service.CustomerService.GetAll()
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	result := gin.H{
		"shipping_costs": shippings,
	}
	helper.SuccessResponseWithData(c, "All shipping costs fetched successfully", http.StatusOK, result)
}

func (shippingHadler *ShippingHadler) ShippingCost(c *gin.Context) {
	var data model.RequestDestination
	if err := c.ShouldBindQuery(&data); err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	cost, err := shippingHadler.Service.CustomerService.ShippingCost(data)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
	}

	result := gin.H{
		"cost": cost,
	}

	helper.SuccessResponseWithData(c, "success", http.StatusOK, result)
}
