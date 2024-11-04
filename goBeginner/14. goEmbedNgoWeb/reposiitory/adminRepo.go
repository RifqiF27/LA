package repository

import (
	"main/model"
)

func (r *UserRepoDb) CreateAdmin(admin *model.Admin) error {
	query := `INSERT INTO "Admin" (name, user_id) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRow(query, admin.Name, admin.UserID).Scan(&admin.ID)
	if err != nil {
		return err
	}
	return nil
}
