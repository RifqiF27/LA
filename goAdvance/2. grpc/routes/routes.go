package routes

import (
	"project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/events", controllers.GetEvents)
		api.GET("/events/:id", controllers.GetEventDetail)
	}
}
