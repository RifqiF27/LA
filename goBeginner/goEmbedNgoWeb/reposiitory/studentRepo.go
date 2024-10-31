package repository

import (
	"database/sql"
	"main/model"
)

type StudentRepository interface {
	AddStudent(student *model.Student) error
	GetAllStudents(students *[]model.Student) error
	UpdateStudent(student *model.Student) error
	DeleteStudent(id, userID int) error
}

type StudentRepoDb struct {
	DB *sql.DB
}

func NewStudentRepo(db *sql.DB) StudentRepository {
	return &StudentRepoDb{DB: db}
}

func (r *StudentRepoDb) AddStudent(student *model.Student) error {
	query := `INSERT INTO "Student" (user_id, name, phone_number, address) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.DB.QueryRow(query, student.UserID, student.Name, student.PhoneNumber, student.Address).Scan(&student.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *StudentRepoDb) GetAllStudents(students *[]model.Student) error {
	query := `SELECT id, user_id, name, phone_number, address FROM "Student"`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var student model.Student
		if err := rows.Scan(&student.ID, &student.UserID, &student.Name, &student.PhoneNumber, &student.Address); err != nil {
			return err
		}
		*students = append(*students, student)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (r *StudentRepoDb) UpdateStudent(student *model.Student) error {
	query := `UPDATE "Student" SET name = $1, phone_number = $2, address = $3 WHERE id = $4 AND user_id = $5`
	_, err := r.DB.Exec(query, student.Name, student.PhoneNumber, student.Address, student.ID, student.UserID)
	return err
}

func (r *StudentRepoDb) DeleteStudent(id, userID int) error {
	query := `DELETE FROM "Student" WHERE id = $1 AND user_id = $2`
	_, err := r.DB.Exec(query, id, userID)
	return err
}
