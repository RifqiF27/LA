package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/database"
	"todo/handler"
	"todo/middleware"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		fmt.Println("Gagal menginisialisasi database:", err)
		return
	}
	defer db.Close()
	serverMux := http.NewServeMux()

	authMux := http.NewServeMux()
	authMux.HandleFunc("/login", handler.Login(db))
	authMux.HandleFunc("/register", handler.Register(db))

	resourceMux := http.NewServeMux()
	resourceMux.HandleFunc("/add-todo", handler.AddTodo(db))
	resourceMux.HandleFunc("/update-todo", handler.UpdateTodoStatus(db))
	resourceMux.HandleFunc("/get-todo", handler.GetTodos(db))

	role := middleware.AdminMiddleware(resourceMux)
	middleware := middleware.Middleware(role)

	serverMux.Handle("/", authMux)
	serverMux.Handle("/customer/", http.StripPrefix("/customer", middleware))


	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", serverMux))

}

// curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "admin", "password": "hashedpassword1"}'
// curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "new_user", "password": "new_password", "role":"staff"}'
// curl -X PUT http://localhost:8080/customer/update-todo -H "Content-Type: application/json" -H "token: admin_token" -d '{"id": 1, "status": "completed"}'
// curl -X POST http://localhost:8080/customer/add-todo -H "Content-Type: application/json" -H "token: admin_token" -d '{"thread": "New Todo Thread", "status": "Processingg"}'
// curl -X GET "http://localhost:8080/customer/get-todo?page=1&limit=5&search=" -H "token: 12345" -H "role: admin_token"-d {}

