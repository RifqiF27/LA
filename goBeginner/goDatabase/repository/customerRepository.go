package repository

import (
	"database/sql"
	"main/model"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
}

type CustomerRepoDb struct {
	TX *sql.Tx
}

func NewCustomerRepo(tx *sql.Tx) CustomerRepository {
	return &CustomerRepoDb{TX: tx}
}

func (r *CustomerRepoDb) Create(customer *model.Customer) error {
	query := `INSERT INTO "Customer" (name, phone_number, address, user_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.TX.QueryRow(query, customer.Name, customer.PhoneNumber, customer.Address, customer.UserID).Scan(&customer.ID)
	if err != nil {
		return err
	}
	return nil
}

