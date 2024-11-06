package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"todo/model"
)

type TodoRepository interface {
	Create(todo *model.Todo) (*model.Todo, error)
	UpdateStatus(status string, id uint16) error
	GetTodosWithPagination(limit, offset int, searchThread string) ([]model.Todo, int, error)
	TodoExists(id uint16) (bool, error)
}

type TodoRepoDb struct {
	DB *sql.DB
}

func NewTodoRepo(db *sql.DB) TodoRepository {
	return &TodoRepoDb{DB: db}
}

func (r *TodoRepoDb) GetTodosWithPagination(limit, offset int, searchThread string) ([]model.Todo, int, error) {
	var todos []model.Todo
	var totalTodos int

	query := `SELECT * FROM "Todos"`
	conditions := []string{}
	args := []interface{}{}

	// if filterStatus {
	// 	conditions = append(conditions, "stock < 10")
	// }

	if searchThread != "" {
		conditions = append(conditions, "thread ILIKE $1")
		args = append(args, searchThread+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += ` ORDER BY "todo_id" ASC`

	if searchThread != "" {
		query += " LIMIT $2 OFFSET $3"
		args = append(args, limit, offset)
	} else {
		query += " LIMIT $1 OFFSET $2"
		args = append(args, limit, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		fmt.Println(err, "<<<")
		return nil, 0, err
	}
	defer rows.Close()
	// fmt.Println(conditions,"<<<<",">>>>", args)
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID,&todo.UserID, &todo.Thread, &todo.Status)
		if err != nil {
			fmt.Println(err, "<<<<<<<")
			return nil, 0, err
		}
		todos = append(todos, todo)
	}

	countQuery := `SELECT COUNT(*) FROM "Todos"`
	if len(conditions) > 0 {
		countQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	var countArgs []interface{}
	if searchThread != "" {
		countArgs = args[:1]
	}
	err = r.DB.QueryRow(countQuery, countArgs...).Scan(&totalTodos)
	if err != nil {
		return nil, 0, err
	}

	return todos, totalTodos, nil
}

func (r *TodoRepoDb) Create(todo *model.Todo) (*model.Todo, error) {
	query := `INSERT INTO "Todos" (user_id, thread, status) VALUES ($1, $2, $3) RETURNING todo_id`
	err := r.DB.QueryRow(query,todo.UserID, todo.Thread, todo.Status).Scan(&todo.ID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepoDb) TodoExists(id uint16) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM "Todos" WHERE id = $1)`

	err := r.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *TodoRepoDb) UpdateStatus(status string, id uint16) error {
	query := `UPDATE "Todos" SET status = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, status, id)
	return err
}
