package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"html/template"
	"todo/model"
	"todo/repository"
	"todo/service"
	"todo/utils"
)

func GetTodos(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Rendering todo")

			err := tmpl.ExecuteTemplate(w, "layout.html",  map[string]bool{"IsTodo": true})
			if err != nil {
				fmt.Println(err, ">>>>>")
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			// w.Header().Get("token")
			// http.Redirect(w, r, "/register", http.StatusSeeOther)
			return

		}
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

		// err = json.NewDecoder(r.Body).Decode(&pagination)
		// if err != nil {
		// 	fmt.Println(err,"ini get body handler")
		// 	utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		// 	return
		// }

		todoService := service.NewTodoService(repository.NewTodoRepo(db))
		todos, totalTodos, err := todoService.GetTodosService(limitInt, pageInt, search)
		if err != nil {
			fmt.Println(err, "ini get handler todo")

			utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		totalPages := int(math.Ceil(float64(totalTodos) / float64(limitInt)))
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

func AddTodo(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
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

		if tmpl != nil {
            tmpl.Execute(w, todo)
        }
	}
}

func UpdateTodoStatus(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
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

		if tmpl != nil {
            tmpl.Execute(w, todo)
        }
	}
}
