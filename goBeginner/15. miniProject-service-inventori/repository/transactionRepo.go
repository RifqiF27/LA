package repository

import (
	"database/sql"
	"fmt"
	"service-inventory/model"
)

type TransactionRepository interface {
	GetTransactions(itemName string) ([]model.TransactionHistory, error)
	CreateTransaction(transaction *model.Transaction) (*model.Transaction, error)
}

type TransactionRepoDb struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) TransactionRepository {
	return &TransactionRepoDb{DB: db}
}

func (r *TransactionRepoDb) GetTransactions(itemName string) ([]model.TransactionHistory, error) {
	var transactions []model.TransactionHistory
	var query string

	if itemName == "" {
		query = `
            SELECT t.transaction_id, i.item_name, t.transaction_type, t.quantity, t.timestamp, t.notes, t.user_id
            FROM "Transactions" t
            JOIN "Items" i ON t.item_id = i.item_id
        `
	} else {
		query = `
            SELECT t.transaction_id, i.item_name, t.transaction_type, t.quantity, t.timestamp, t.notes, t.user_id
            FROM "Transactions" t
            JOIN "Items" i ON t.item_id = i.item_id
            WHERE i.item_name ILIKE $1
        `
	}

	var rows *sql.Rows
	var err error
	if itemName == "" {
		rows, err = r.DB.Query(query)
	} else {
		rows, err = r.DB.Query(query, itemName+"%")
	}

	if err != nil {
		fmt.Println(err, "<<<<<< rows")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.TransactionHistory

		err := rows.Scan(&transaction.TransactionID, &transaction.ItemName, &transaction.TransactionType, &transaction.Quantity, &transaction.Timestamp, &transaction.Notes, &transaction.UserID)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
		fmt.Println("Transaction added:", transaction)
	}

	if len(transactions) == 0 {
		fmt.Println("No transactions found.")
	} else {
		fmt.Printf("%d transactions found.\n", len(transactions))
	}

	return transactions, nil
}

func (r *TransactionRepoDb) CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	query := `
        INSERT INTO "Transactions" (item_id, transaction_type, quantity, timestamp, notes, user_id)
        VALUES ($1, $2, $3, NOW(), $4, $5)
        RETURNING transaction_id, timestamp
    `
	
	err := r.DB.QueryRow(query, transaction.ItemID, transaction.TransactionType, transaction.Quantity, transaction.Notes, transaction.UserID).
		Scan(&transaction.TransactionID, &transaction.Timestamp)
	if err != nil {
		return nil, err
	}

	var updateQuery string
	if transaction.TransactionType == "in" {
		updateQuery = `UPDATE "Items" SET stock = stock + $1 WHERE item_id = $2`
	} else if transaction.TransactionType == "out" {
		updateQuery = `UPDATE "Items" SET stock = stock - $1 WHERE item_id = $2`
	}

	_, err = r.DB.Exec(updateQuery, transaction.Quantity, transaction.ItemID)
	if err != nil {
		return nil, err
	}
	
	return transaction, nil
}
