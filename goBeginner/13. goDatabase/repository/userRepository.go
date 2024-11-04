package repository

import (
	"database/sql"
	"main/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetAll(user *[]model.User) error
	GetUserActive(user *[]model.UserLog) error
}
type UserRepoDb struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &UserRepoDb{DB: db}
}

func (r *UserRepoDb) Create(user *model.User) error {
	query := `INSERT INTO "User" (username, password, role) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, user.Username, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepoDb) GetAll(users *[]model.User) error {
	query := `SELECT id, username, password, role FROM "User" WHERE role != 'admin'`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role); err != nil {
			return err
		}
		*users = append(*users, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
func (r *UserRepoDb) GetUserActive(users *[]model.UserLog) error {
	query := `
    SELECT
    (
        SELECT
            COUNT(*)
        FROM
            "Login_Log"
        WHERE
            "status" = true
    ) AS "active_customers",
    (
        SELECT
            COUNT(*)
        FROM
            "Login_Log"
        WHERE
            "status" = false
    ) AS "inactive_customers"
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.UserLog
		if err := rows.Scan(&user.ActiveCustomer, &user.InactiveCustomer); err != nil {
			return err
		}
		*users = append(*users, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
