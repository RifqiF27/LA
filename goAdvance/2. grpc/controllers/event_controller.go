package controllers

import (
	"net/http"
	"project/models"
	"project/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	utils.JSONResponse(c, http.StatusOK, "success", models.Events)
}

func GetEventDetail(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "invalid event ID", nil)
		return
	}

	for _, event := range models.Events {
		if event.ID == id {
			utils.JSONResponse(c, http.StatusOK, "success", event)
			return
		}
	}

	utils.JSONResponse(c, http.StatusNotFound, "event not found", nil)
}
