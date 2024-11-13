package repository

import (
	"database/sql"
	"fmt"
	"log"
	"travelika/model"

	"go.uber.org/zap"
)

type DestinationRepository interface {
	GetAllEvents(eventName, location string, date string, orderBy string, orderAsc bool, limit, page int) ([]model.DestinationEventRating, int, int, error)
}

type destinationRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewDestinationRepository(db *sql.DB, logger *zap.Logger) DestinationRepository {
	return &destinationRepository{db: db, log: logger}
}

func (r *destinationRepository) GetAllEvents(eventName, location string, date string, orderBy string, orderAsc bool, limit, page int) ([]model.DestinationEventRating, int, int, error) {

	query := `SELECT d.id, d.location, d.image_url, d.description as destination, 
                     e.name as event, e.schedule, e.price, 
                     COALESCE(AVG(r.rating), 0) as average_rating
              FROM destinations d
              JOIN events e ON d.id = e.destination_id
              LEFT JOIN reviews r ON e.id = r.destination_id
              WHERE 1=1`

	countQuery := `SELECT COUNT(*) 
                   FROM destinations d 
                   JOIN events e ON d.id = e.destination_id 
                   LEFT JOIN reviews r ON e.id = r.destination_id 
                   WHERE 1=1`

	var params []interface{}
	paramIndex := 1

	if eventName != "" {
		query += ` AND e.name ILIKE $` + fmt.Sprint(paramIndex)
		countQuery += ` AND e.name ILIKE $` + fmt.Sprint(paramIndex)
		params = append(params, "%"+eventName+"%")
		paramIndex++
	}

	if location != "" {
		query += ` AND d.location ILIKE $` + fmt.Sprint(paramIndex)
		countQuery += ` AND d.location ILIKE $` + fmt.Sprint(paramIndex)
		params = append(params, "%"+location+"%")
		paramIndex++
	}

	if date != "" {
		query += ` AND e.schedule = $` + fmt.Sprint(paramIndex)
		countQuery += ` AND e.schedule = $` + fmt.Sprint(paramIndex)
		params = append(params, date)
		paramIndex++
	}

	orderColumn := "d.id"
	switch orderBy {
	case "name":
		orderColumn = "e.name"
	case "location":
		orderColumn = "d.location"
	case "price":
		orderColumn = "e.price"
	}

	orderDirection := "DESC"
	if !orderAsc {
		orderDirection = "ASC"
	}

	query += ` GROUP BY d.id, d.location, d.image_url, d.description, e.name, e.schedule, e.price`
	query += ` ORDER BY ` + orderColumn + ` ` + orderDirection

	var totalItems int
	err := r.db.QueryRow(countQuery, params...).Scan(&totalItems)
	if err != nil {
		r.log.Error("Repository: failed to execute count query", zap.Error(err))
		return nil, 0, 0, err
	}

	totalPages := (totalItems + limit - 1) / limit
	offset := (page - 1) * limit
	query += ` LIMIT $` + fmt.Sprint(paramIndex) + ` OFFSET $` + fmt.Sprint(paramIndex+1)
	params = append(params, limit, offset)

	rows, err := r.db.Query(query, params...)
	if err != nil {
		r.log.Error("Repository: failed to execute query", zap.Error(err))
		return nil, 0, 0, err
	}
	defer rows.Close()

	r.log.Info("Repository: executed query", zap.String("query", query), zap.Any("params", params))

	log.Printf("Query: %s", query)
	log.Printf("Params: %v", params)

	var results []model.DestinationEventRating
	for rows.Next() {
		var result model.DestinationEventRating
		if err := rows.Scan(&result.ID, &result.Location, &result.ImageURL, &result.Destination, &result.EventName,
			&result.Schedule, &result.Price, &result.AverageRating); err != nil {
			r.log.Error("Repository: failed to scan row", zap.Error(err))
			return nil, 0, 0, err
		}
		results = append(results, result)
	}

	return results, totalItems, totalPages, nil
}
