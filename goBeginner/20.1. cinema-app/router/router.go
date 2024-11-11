package router

import (
	"cinema/database"
	"cinema/handler"
	middleware_auth "cinema/middleware"
	"cinema/repository"
	"cinema/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	db := database.NewPostgresDB()

	userRepo := repository.NewAuthRepositoryDb(db)
	srv := service.NewAuthService(userRepo)

	cinemaRepo := repository.NewCinemaRepositoryDb(db)
	srvCnms := service.NewCinemaService(cinemaRepo)

	h := handler.NewAuthHandler(srv)
	c := handler.NewCinemaHandler(srvCnms)

	r.Use(middleware.Logger)

	// Serve static files
	// fileServer := http.FileServer(http.Dir("./static"))
	// r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Post("/register", h.RegisterHandler)
			r.Post("/login", h.LoginHandler)
			// r.Post("/logout", h.LogoutHandler)

		})
	})

	r.Group(func(r chi.Router) {
		r.Route("/api/cinemas", func(r chi.Router) {
			r.Use(middleware_auth.AuthMiddleware(srv)) // Apply the AuthMiddleware
			r.Get("/", c.GetAllCinemas)
			r.Get("/{cinemaId}", c.GetCinemaByID)
			r.Get("/{cinemaId}/seats", c.GetSeats)

			r.Post("/logout", h.LogoutHandler)

		})
	})

	return r
}
