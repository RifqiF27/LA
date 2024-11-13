package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
	"travelika/handler"
)

func NewRouter(authHandler *handler.AuthHandler, DestinationHandler *handler.DestinationHandler, log *zap.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			log.Info("Request received", zap.String("method", r.Method), zap.String("url", r.URL.String()))

			next.ServeHTTP(w, r)

			duration := time.Since(start)

			log.Info("Request processed",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.Duration("duration", duration),
			)
		})
	})

	fileServer := http.FileServer(http.Dir("./uploads"))
	r.Handle("/uploads/*", http.StripPrefix("/uploads", fileServer))

	r.Group(func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)
		r.Post("/logout", authHandler.Logout)
	})

	r.Group(func(r chi.Router) {
		r.Route("/api/destinations", func(r chi.Router) {
			r.Get("/", DestinationHandler.GetDestination)
		})

	})

	return r
}
