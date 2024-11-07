package main

import (
	"fmt"
	"html/template"
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

	// tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl, err := template.ParseGlob("tmpl/templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	// tmpl := template.Must(template.ParseFiles(
	// 	"templates/layout.html",
	// 	"templates/header.html",
	// 	"templates/login.html",
	// 	"templates/navigator.html",
	// 	"templates/content.html",
	// 	"templates/footer.html",
	// ))

	// fs := http.FileServer(http.Dir("tmpl/static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	serverMux := http.NewServeMux()

	authMux := http.NewServeMux()
	authMux.HandleFunc("/login", handler.Login(db, tmpl))
	authMux.HandleFunc("/register", handler.Register(db, tmpl))

	resourceMux := http.NewServeMux()
	resourceMux.HandleFunc("/add-todo", handler.AddTodo(db, tmpl))
	resourceMux.HandleFunc("/update-todo", handler.UpdateTodoStatus(db, tmpl))
	resourceMux.HandleFunc("/get-todo", handler.GetTodos(db, tmpl))

	role := middleware.AdminMiddleware(resourceMux)
	middleware := middleware.Middleware(role)

	serverMux.Handle("/", authMux)
	serverMux.Handle("/todo/", http.StripPrefix("/todo", middleware))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", serverMux))

}

// curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "admin", "password": "hashedpassword1"}'
// curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "new_user", "password": "new_password", "role":"staff"}'
// curl -X PUT http://localhost:8080/todo/update-todo -H "Content-Type: application/json" -H "token: admin_token" -d '{"id": 1, "status": "completed"}'
// curl -X POST http://localhost:8080/todo/add-todo -H "Content-Type: application/json" -H "token: admin_token" -d '{"thread": "New Todo Thread", "status": "Processingg"}'
// curl -X GET "http://localhost:8080/todo/get-todo?page=1&limit=5&search=" -H "role: admin_token" -H "token: 12345" -d {}
