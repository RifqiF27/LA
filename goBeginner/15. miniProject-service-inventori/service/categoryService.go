package service

import (
	"errors"
	"service-inventory/model"
	"service-inventory/repository"
)

type CategoryService struct {
	RepoCategory repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{RepoCategory: repo}
}

func (cs *CategoryService) AddCategoryService( name string) error {
	if name == "" {
		return errors.New("Category name cannot be empty")
	}
	

	category := model.Category{
		CategoryName: name,
	}
	_, err := cs.RepoCategory.Create(&category)
	if err != nil {
		return errors.New("failed to create category: " + err.Error())
	}

	return nil
}