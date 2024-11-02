package handler

import (
	"database/sql"
	"service-inventory/model"
	"service-inventory/repository"
	"service-inventory/service"
	"service-inventory/utils"
)

func AddLocation(db *sql.DB) {

	_, valid := utils.SessionAdmin()
	if !valid {
		return
	}

	location := model.Location{}
	err := utils.DecodeJSONFile("body.json", &location)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}
	repo := repository.NewLocationRepo(db)
	locationService := service.NewLocationService(repo)

	err = locationService.AddLocationService(location.LocationName)
	if err != nil {
		if err.Error() == "Location name cannot be empty" {
			utils.SendJSONResponse(400, err.Error(), nil)
		} else {
			utils.SendJSONResponse(500, err.Error(), nil)
		}
		return
	}

	utils.SendJSONResponse(201, "location added successfully", location)
}
