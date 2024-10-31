package repository

import (
	"main/model"
)

func (r *UserRepoDb) CreateAdmin(admin *model.Admin) error {
    query := `INSERT INTO "Admin" (name, user_id) VALUES ($1, $2) RETURNING id`
    err := r.DB.QueryRow(query, admin.Name, admin.UserID,).Scan(&admin.ID)
    if err != nil {
        return err
    }
    return nil
}

// type AdminRepositoryDB struct {
// 	DB *sql.DB
// }

// func NewAdminRepository(db *sql.DB) AdminRepositoryDB {
// 	return AdminRepositoryDB{DB: db}
// }

// func (r *AdminRepositoryDB) Create(user *model.Admin) error {
// 	query := `INSERT INTO Admin (name, user_id, class_id) VALUES ($1, $2, $3) RETURNING id`
// 	err := r.DB.QueryRow(query, user.Name, user.UserID, user.ClassID).Scan(&user.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *AdminRepositoryDB) GetAll() (*[]model.Admin, error) {
// 	query := `SELECT id, name, user_id, class_id FROM "Admin"`
// 	rows, err := r.DB.Query(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	admins := []model.Admin{}

// 	for rows.Next() {
// 		var admin model.Admin
// 		rows.Scan(&admin.ID, &admin.Name, &admin.UserID, &admin.ClassID)

// 		admins = append(admins, admin)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return &admins, nil
// }

