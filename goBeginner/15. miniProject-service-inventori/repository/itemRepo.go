package repository

import (
	"database/sql"
	"service-inventory/model"
	"strings"
)

type ItemRepository interface {
	Create(item *model.Item) (*model.Item, error)
	UpdateStock(itemCode string, stock int) error
	GetItemsWithPagination(limit, offset int, filterStock bool, searchName string) ([]model.Item, int, error)
	ItemExists(itemCode string) (bool, error)
}

type ItemRepoDb struct {
	DB *sql.DB
}

func NewItemRepo(db *sql.DB) ItemRepository {
	return &ItemRepoDb{DB: db}
}

func (r *ItemRepoDb) GetItemsWithPagination(limit, offset int, filterStock bool, searchName string) ([]model.Item, int, error) {
	var items []model.Item
	var totalItems int

	query := `SELECT * FROM "Items"`
	conditions := []string{}
	args := []interface{}{}

	if filterStock {
		conditions = append(conditions, "stock < 10")
	}

	if searchName != "" {
		conditions = append(conditions, "item_name ILIKE $1")
		args = append(args, searchName+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += ` ORDER BY "item_name" ASC`

	if searchName != "" {
		query += " LIMIT $2 OFFSET $3"
		args = append(args, limit, offset)
	} else {
		query += " LIMIT $1 OFFSET $2"
		args = append(args, limit, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.ItemCode, &item.ItemName, &item.Stock, &item.CategoryID, &item.LocationID)
		if err != nil {
			return nil, 0, err
		}
		items = append(items, item)
	}

	countQuery := `SELECT COUNT(*) FROM "Items"`
	if len(conditions) > 0 {
		countQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	var countArgs []interface{}
	if searchName != "" {
		countArgs = args[:1]
	}
	err = r.DB.QueryRow(countQuery, countArgs...).Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	return items, totalItems, nil
}

func (r *ItemRepoDb) Create(item *model.Item) (*model.Item, error) {
	query := `INSERT INTO "Items" (item_code, item_name, stock, category_id, location_id) VALUES ($1, $2, $3, $4, $5) RETURNING item_id`
	err := r.DB.QueryRow(query, item.ItemCode, item.ItemName, item.Stock, item.CategoryID, item.LocationID).Scan(&item.ID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *ItemRepoDb) ItemExists(itemCode string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM "Items" WHERE item_code = $1)`

	err := r.DB.QueryRow(query, itemCode).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *ItemRepoDb) UpdateStock(itemCode string, stock int) error {
	query := `UPDATE "Items" SET stock = $1 WHERE item_code = $2`
	_, err := r.DB.Exec(query, stock, itemCode)
	return err
}
