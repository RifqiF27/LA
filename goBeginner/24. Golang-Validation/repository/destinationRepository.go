package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"travelika/model"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type DestinationRepository interface {
	GetAllEvents(eventName, location string, date string, orderBy string, orderAsc bool, limit, page int) ([]model.DestinationEventRating, int, int, error)
	GetById(id int) (*model.DestinationEventRating, error)
	GetTourPlansByEventID(eventID int) (*model.TourPlan, error)
	GetLocationsByDestinationID(destinationID int) (*model.Location, error)
}

type destinationRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewDestinationRepository(db *sql.DB, logger *zap.Logger) DestinationRepository {
	return &destinationRepository{db: db, log: logger}
}

func (r *destinationRepository) GetAllEvents(eventName, location string, date string, orderBy string, orderAsc bool, limit, page int) ([]model.DestinationEventRating, int, int, error) {

	query := `
			SELECT d.id, d.location, d.image_url, d.description, 
			e.name as event, TO_CHAR(e.schedule, 'yyyy-mm-dd'), e.price, 
			COALESCE(AVG(r.rating), 0) as average_rating,
			COUNT(CASE WHEN t.status = TRUE THEN 1 END) AS people
			FROM destinations d
			JOIN events e ON d.id = e.destination_id
			LEFT JOIN reviews r ON e.id = r.destination_id
			LEFT JOIN transactions t ON e.id = t.event_id
			WHERE 1=1
			`

	countQuery := `
				SELECT COUNT(*) 
                FROM destinations d 
                JOIN events e ON d.id = e.destination_id 
                LEFT JOIN reviews r ON e.id = r.destination_id 
                WHERE 1=1
				`

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
		if err := rows.Scan(&result.ID, &result.Location, &result.ImageURL, &result.Description, &result.EventName,
			&result.Schedule, &result.Price, &result.AverageRating, &result.People); err != nil {
			r.log.Error("Repository: failed to scan row", zap.Error(err))
			return nil, 0, 0, err
		}
		results = append(results, result)
	}

	return results, totalItems, totalPages, nil
}

func (r *destinationRepository) GetById(id int) (*model.DestinationEventRating, error) {
	var destination model.DestinationEventRating
	var galleryID []int64
	var galleryImages []string
	var galleryDescriptions []string

	query := `
	SELECT d.id, d.location, d.image_url, d.description,
    array_agg(DISTINCT g.id) AS gallery_id, 
    array_agg(DISTINCT g.image_url) AS gallery_images, 
    array_agg(DISTINCT g.description) AS gallery_descriptions,
    e.name AS event, TO_CHAR(e.schedule, 'yyyy-mm-dd'), e.price, 
    COALESCE(AVG(r.rating), 0) AS average_rating,
    COUNT(DISTINCT CASE WHEN t.status = TRUE THEN t.id END) AS people
	FROM destinations d
	JOIN events e ON d.id = e.destination_id
	LEFT JOIN reviews r ON e.id = r.destination_id
	LEFT JOIN gallery g ON d.id = g.destination_id
	LEFT JOIN transactions t ON e.id = t.event_id
	WHERE d.id = $1
	GROUP BY d.id, d.location, d.image_url, d.description, e.id, e.name, e.schedule, e.price
	`
	err := r.db.QueryRow(query, id).Scan(&destination.ID, &destination.Location, &destination.ImageURL, &destination.Description, pq.Array(&galleryID), pq.Array(&galleryImages), pq.Array(&galleryDescriptions), &destination.EventName, &destination.Schedule, &destination.Price, &destination.AverageRating, &destination.People)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.log.Error("Repository: failed to scan row", zap.Error(err))
		return nil, err
	}
	for i := 0; i < len(galleryImages); i++ {
		destination.Gallery = append(destination.Gallery, model.Gallery{
			ID:          int(galleryID[i]),
			ImageURL:    galleryImages[i],
			Description: galleryDescriptions[i],
		})
	}
	return &destination, nil
}

func (r *destinationRepository) GetTourPlansByEventID(eventID int) (*model.TourPlan, error) {
	query := `SELECT destination_id, event_id, day, activity, facilities FROM tour_plans WHERE event_id = $1`
	rows, err := r.db.Query(query, eventID)
	if err != nil {
		r.log.Error("Repository: ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	tourPlansMap := make(map[int]*model.TourPlan)

	for rows.Next() {
		var day int
		var activity string
		var facilityJSON string
		plan := model.TourPlan{}

		if err := rows.Scan(&plan.DestinationID, &plan.EventID, &day, &activity, &facilityJSON); err != nil {
			r.log.Error("Repository: failed to scan row", zap.Error(err))
			return nil, err
		}

		var facility model.Facility
		if err := json.Unmarshal([]byte(facilityJSON), &facility); err != nil {
			r.log.Error("Repository: failed to unmarshal json facilities", zap.Error(err))
			return nil, err
		}

		if _, exists := tourPlansMap[plan.DestinationID]; !exists {
			plan.Days = []model.DayPlan{}
			tourPlansMap[plan.DestinationID] = &plan
		}

		tourPlansMap[plan.DestinationID].Days = append(tourPlansMap[plan.DestinationID].Days, model.DayPlan{
			Day:      day,
			Activity: activity,
			Facility: facility,
		})
	}

	if err := rows.Err(); err != nil {
		r.log.Error("Repository: failed to scan row", zap.Error(err))
		return nil, err
	}

	for _, plan := range tourPlansMap {
		return plan, nil
	}

	return nil, nil
}

func (r *destinationRepository) GetLocationsByDestinationID(destinationID int) (*model.Location, error) {
	var loc model.Location

	query := `
        SELECT id, destination_id, summary, longlat[0] AS longitude, longlat[1] AS latitude, detail
        FROM locations
        WHERE destination_id = $1
    `

	err := r.db.QueryRow(query, destinationID).Scan(&loc.ID, &loc.DestinationID, &loc.Summary, &loc.Longitude, &loc.Latitude, &loc.Detail)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.log.Error("Repository: failed to scan row", zap.Error(err))
		return nil, err
	}

	return &loc, nil
}
