package handler

import (
	"database/sql"
	"service-inventory/model"
	"service-inventory/repository"
	"service-inventory/service"
	"service-inventory/utils"
)

func AddCategory(db *sql.DB) {

	_, valid := utils.SessionAdmin()
	if !valid {
		return
	}

	category := model.Category{}
	err := utils.DecodeJSONFile("body.json", &category)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}
	repo := repository.NewCategoryRepo(db)
	categoryService := service.NewCategoryService(repo)

	err = categoryService.AddCategoryService(category.CategoryName)
	if err != nil {
		if err.Error() == "Category name cannot be empty" {
			utils.SendJSONResponse(400, err.Error(), nil)
		} else {
			utils.SendJSONResponse(500, err.Error(), nil)
		}
		return
	}

	utils.SendJSONResponse(201, "category added successfully", category)
}
