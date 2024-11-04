package repository

import (
	"database/sql"
	"service-inventory/model"
)

type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	
}

type CategoryRepoDb struct {
	DB *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepository {
	return &CategoryRepoDb{DB: db}
}


func (r *CategoryRepoDb) Create(category *model.Category) (*model.Category, error) {
	query := `INSERT INTO "Categories" (category_name) VALUES ($1) RETURNING category_id`
	err := r.DB.QueryRow(query, category.CategoryName).Scan(&category.ID)
	if err != nil {
		return nil, err
	}

	return category, nil
}


