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
	customer := v1.Group("/shipping")
	{
		customer.Use(middleware.Authentication())
		customer.GET("/", ctx.Handler.ShippingHandler.GetAllShipping)
		customer.GET("/cost", ctx.Handler.ShippingHandler.ShippingCost)
	}

	fmt.Println("server start on port ", ctx.Config.Port)
	router.Run(":" + ctx.Config.Port)
}
