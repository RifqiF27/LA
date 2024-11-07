package router

import (
	"book-store/database"
	"book-store/handler"
	"book-store/middleware"
	"book-store/repository"
	"book-store/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	db := database.NewPostgresDB()

	repo := repository.NewUserRepo(db)
	srv := service.NewUserService(repo)
	h := handler.NewAuthHandler(srv)

	

	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Group(func(r chi.Router) {
		r.Get("/login", h.Login)
		r.Post("/login", h.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware_auth.ValidateJWT)
		r.Get("/dashboard", h.Dashboard)
		// r.Get("/products/{id}", h.GetProductByID)
		// r.Put("/products/{id}", h.UpdateProduct)
		// r.Delete("/products/{id}", h.DeleteProduct)

	})

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return cors(r)
}
