package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"service-inventory/model"
	"service-inventory/repository"
	"service-inventory/service"
	"service-inventory/utils"
)

func GetItems(db *sql.DB) {
	_, valid := utils.Session()
	if !valid {
		return
	}
	var pagination model.PaginationRequest

	err := utils.DecodeJSONFile("body.json", &pagination)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	itemService := service.NewItemService(repository.NewItemRepo(db))
	items, totalItems, err := itemService.GetItemsService(pagination.Limit, pagination.Page, pagination.FilterStock, pagination.SearchName)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))

	response := model.Response{
		StatusCode: 200,
		Message:    "Data retrieved successfully",
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       items,
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

}

func AddItem(db *sql.DB) {

	_, valid := utils.SessionAdmin()
	if !valid {
		return
	}

	item := model.Item{}
	err := utils.DecodeJSONFile("body.json", &item)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}
	repo := repository.NewItemRepo(db)
	itemService := service.NewItemService(repo)

	err = itemService.AddItemService(item.ItemCode, item.ItemName, item.Stock, item.CategoryID, item.LocationID)
	if err != nil {
		if err.Error() != "failed to create item: pq: duplicate key value violates unique constraint \"Items_item_code_key\"" {
			utils.SendJSONResponse(400, err.Error(), nil)
		} else {
			utils.SendJSONResponse(500, err.Error(), nil)
		}
		return
	}

	utils.SendJSONResponse(201, "item added successfully", item)
}

func UpdateItemStock(db *sql.DB) {

	_, valid := utils.SessionAdmin()
	if !valid {
		return
	}

	item := model.Item{}

	err := utils.DecodeJSONFile("body.json", &item)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	repo := repository.NewItemRepo(db)
	itemService := service.NewItemService(repo)

	err = itemService.UpdateStockService(item.ItemCode, item.Stock)
	if err != nil {
		if err.Error() == "item code cannot be empty" || err.Error() == "stock cannot be negative" {
			utils.SendJSONResponse(400, err.Error(), nil)
		} else if err.Error() == "item not found" {
			utils.SendJSONResponse(404, err.Error(), nil)
		} else {
			utils.SendJSONResponse(500, err.Error(), nil)
		}
		return
	}

	utils.SendJSONResponse(200, "stock updated successfully", item)
}
