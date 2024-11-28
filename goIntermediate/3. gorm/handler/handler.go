package handler

import (
	"ecommerce/service"

	"go.uber.org/zap"
)

type AllHandler struct {
	ShippingHandler ShippingHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger) AllHandler {
	return AllHandler{
		ShippingHandler: NewShippingHandler(service, log),
	}
}
