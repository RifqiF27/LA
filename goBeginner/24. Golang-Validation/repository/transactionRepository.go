package repository

import (
	"database/sql"
	"travelika/model"

	"go.uber.org/zap"
)

type TransactionsRepository interface {
	Create(transaction model.Transaction) error
}

type transactionRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewTransactionRepository(db *sql.DB, logger *zap.Logger) TransactionsRepository  {
	return &transactionRepository{db: db, log: logger}
}

func (r *transactionRepository) Create(transaction model.Transaction) error {
	query := `INSERT INTO transactions (name, email, phone, comment, event_id, status) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int
	err := r.db.QueryRow(query, transaction.Name, transaction.Email, transaction.Phone, transaction.Comment, transaction.EventID, transaction.StatusTrx).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}