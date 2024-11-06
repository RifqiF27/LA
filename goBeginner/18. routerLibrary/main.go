package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"todo/database"
	"todo/handler"
	"todo/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	// serverMux := http.NewServeMux()

	// authMux := http.NewServeMux()
	// authMux.HandleFunc("/login", handler.Login(db, tmpl))
	// authMux.HandleFunc("/register", handler.Register(db, tmpl))

	// resourceMux := http.NewServeMux()
	// resourceMux.HandleFunc("/add-todo", handler.AddTodo(db, tmpl))
	// resourceMux.HandleFunc("/update-todo", handler.UpdateTodoStatus(db, tmpl))
	// resourceMux.HandleFunc("/get-todo", handler.GetTodos(db, tmpl))

	// role := middleware.AdminMiddleware(resourceMux)
	// middleware := middleware.Middleware(role)

	// serverMux.Handle("/", authMux)
	// serverMux.Handle("/todo/", http.StripPrefix("/todo", middleware))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Get("/login", handler.Login(db, tmpl))
		r.Post("/login", handler.Login(db, tmpl))
		r.Get("/register", handler.Register(db, tmpl))
		r.Post("/register", handler.Register(db, tmpl))
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware_test.AdminMiddleware)
		r.Get("/todo/get-todo", handler.GetTodos(db, tmpl))
		r.Post("/todo/add-todo", handler.AddTodo(db, tmpl))
		r.Put("/todo/update-todo", handler.UpdateTodoStatus(db, tmpl))
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
