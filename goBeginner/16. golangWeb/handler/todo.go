package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"todo/model"
	"todo/repository"
	"todo/service"
	"todo/utils"
)

func GetTodos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var pagination model.PaginationRequest
		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")
		search := r.URL.Query().Get("search")

		pageInt := 1
		limitInt := 10
		var err error

		if page != "" {
			pageInt, err = strconv.Atoi(page)
			if err != nil || pageInt < 1 {
				pageInt = 1
			}
		}

		if limit != "" {
			limitInt, err = strconv.Atoi(limit)
			if err != nil || limitInt < 1 {
				limitInt = 10
			}
		}

		err = json.NewDecoder(r.Body).Decode(&pagination)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		todoService := service.NewTodoService(repository.NewTodoRepo(db))
		todos, totalTodos, err := todoService.GetTodosService(limitInt, pageInt, search)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		totalPages := int(math.Ceil(float64(totalTodos) / float64(pagination.Limit)))

		response := model.Response{
			StatusCode: http.StatusOK,
			Message:    "Data retrieved successfully",
			Page:       pagination.Page,
			Limit:      pagination.Limit,
			TotalTodos: totalTodos,
			TotalPages: totalPages,
			Data:       todos,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func AddTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("AddTodo Handler Invoked")
		todo := model.Todo{}
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		repo := repository.NewTodoRepo(db)
		todoService := service.NewTodoService(repo)

		err = todoService.AddTodoService(todo.UserID, todo.Thread, todo.Status)
		if err != nil {
			if err.Error() != "failed to create todo: pq: duplicate key value violates unique constraint \"Todos_todo_code_key\"" {
				utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			} else {
				utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}

		utils.SendJSONResponse(w, http.StatusCreated, "Todo added successfully", todo)
	}
}

func UpdateTodoStatus(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		todo := model.Todo{}
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		repo := repository.NewTodoRepo(db)
		todoService := service.NewTodoService(repo)

		err = todoService.UpdateStatusService(todo.ID, todo.Status)
		if err != nil {
			switch err.Error() {
			case "todo code cannot be empty", "stock cannot be negative":
				utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			case "todo not found":
				utils.SendJSONResponse(w, http.StatusNotFound, err.Error(), nil)
			default:
				utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}

		utils.SendJSONResponse(w, http.StatusOK, "Stock updated successfully", todo)
	}
}
