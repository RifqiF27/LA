package service

import (
	"errors"
	"service-inventory/model"
	"service-inventory/repository"
)

type ItemService struct {
	RepoItem repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) *ItemService {
	return &ItemService{RepoItem: repo}
}

func (is *ItemService) GetItemsService(limit, page int, filterStock bool, searchName string) ([]model.Item, int, error) {
	offset := (page - 1) * limit

	items, totalItems, err := is.RepoItem.GetItemsWithPagination(limit, offset, filterStock, searchName)
	if err != nil {
		return nil, 0, errors.New("failed to retrieve items")
	}

	return items, totalItems, nil
}

func (is *ItemService) AddItemService(code, name string, stock, categoryID, locationID int) error {
	if name == "" {
		return errors.New("item name cannot be empty")
	}
	if code == "" {
		return errors.New("item code cannot be empty")
	}
	if stock <= 0 {
		return errors.New("stock cannot be negative or 0")
	}
	if categoryID <= 0 {
		return errors.New("category id cannot be negative or 0")
	}
	if locationID <= 0 {
		return errors.New("location id cannot be negative or 0")
	}

	item := model.Item{
		ItemCode:   code,
		ItemName:   name,
		Stock:      stock,
		CategoryID: categoryID,
		LocationID: locationID,
	}
	_, err := is.RepoItem.Create(&item)
	if err != nil {
		return errors.New("failed to create item: " + err.Error())
	}

	// fmt.Println("Successfully added item with ID:", item.ID)
	return nil
}

func (is *ItemService) UpdateStockService(itemCode string, newStock int) error {

	if itemCode == "" {
		return errors.New("item code cannot be empty")
	}
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}

	exists, err := is.RepoItem.ItemExists(itemCode)
	if err != nil {
		return errors.New("failed to check item existence: " + err.Error())
	}
	if !exists {
		return errors.New("item not found")
	}

	err = is.RepoItem.UpdateStock(itemCode, newStock)
	if err != nil {
		return errors.New("failed to update stock: " + err.Error())
	}

	return nil
}
