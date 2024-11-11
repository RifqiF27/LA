package repository

import (
	"cinema/model"
	"database/sql"
)

type CinemaRepository interface {
	GetAll() ([]model.Cinema, error)
	GetByID(id int) (*model.Cinema, error)
	GetSeats(cinemaID int, date, time string) ([]model.Seat, error)
}

type CinemaRepositoryDb struct {
	DB *sql.DB
}

func NewCinemaRepositoryDb(db *sql.DB) CinemaRepository {
	return &CinemaRepositoryDb{DB: db}
}

func (r *CinemaRepositoryDb) GetAll() ([]model.Cinema, error) {
	query := `SELECT id, name, location FROM cinemas`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cinemas []model.Cinema
	for rows.Next() {
		var cinema model.Cinema
		if err := rows.Scan(&cinema.ID, &cinema.Name, &cinema.Location); err != nil {
			return nil, err
		}
		cinemas = append(cinemas, cinema)
	}

	return cinemas, nil
}

func (r *CinemaRepositoryDb) GetByID(id int) (*model.Cinema, error) {
	query := `SELECT 
    c.id, 
    c.name, 
    c.location,
    COALESCE(COUNT(s.id), 0) AS seats
FROM 
    cinemas c
JOIN 
    seats s ON c.id = s.cinema_id
WHERE 
    c.id = $1
GROUP BY 
    c.id, c.name, c.location`
	var cinema model.Cinema
	err := r.DB.QueryRow(query, id).Scan(&cinema.ID, &cinema.Name, &cinema.Location, &cinema.Seats)
	if err != nil {
		return nil, err
	}
	return &cinema, nil
}

func (r *CinemaRepositoryDb) GetSeats(cinemaID int, date, time string) ([]model.Seat, error) {
	query := `SELECT id, cinema_id, date, time, status FROM seats WHERE cinema_id = $1 AND date = $2 AND time = $3`
	rows, err := r.DB.Query(query, cinemaID, date, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var seats []model.Seat
	for rows.Next() {
		var seat model.Seat
		err := rows.Scan(&seat.ID, &seat.CinemaID, &seat.Date, &seat.Time, &seat.Status)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}
