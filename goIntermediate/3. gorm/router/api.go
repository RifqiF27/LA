package router

import (
	"ecommerce/infra"
	"ecommerce/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupReouter(ctx infra.Context) {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.Use(middleware.Logger())

	// shipping routes
	r := v1.Group("/shipping")
	{
		r.Use(middleware.Authentication())
		r.GET("/", ctx.Handler.ShippingHandler.GetAllShipping)
		r.GET("/cost", ctx.Handler.ShippingHandler.ShippingCost)
		r.POST("/", ctx.Handler.ShippingHandler.CreateNewShippingHandler)
		r.GET("/:order_shipping_id", ctx.Handler.ShippingHandler.TrackDeliveryHandler)
		r.PUT("/:order_shipping_id/status", ctx.Handler.ShippingHandler.UpdateShippingStatusHandler)
	}

	fmt.Println("server start on port ", ctx.Config.Port)
	router.Run(":" + ctx.Config.Port)
}
