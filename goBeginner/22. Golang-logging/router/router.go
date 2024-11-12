package router

import (
	"book-store/handler"
	middleware_auth "book-store/middleware"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"go.uber.org/zap"
)

func NewRouter(authHandler *handler.AuthHandler, bookHandler *handler.BookHandler, paymentHandler *handler.PaymentMethodHandler, log *zap.Logger) http.Handler {
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

	fileServer := http.FileServer(http.Dir("./assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	r.Group(func(r chi.Router) {
		r.Get("/login", authHandler.Login)
		r.Post("/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware_auth.ValidateJWT(log))
		r.Route("/books", func(r chi.Router) {
			r.Get("/", bookHandler.GetAllBooks)
			r.Post("/", bookHandler.CreateBook)
			r.Put("/", bookHandler.UpdateBook)
			r.Get("/{id}", bookHandler.GetBookByID)
			r.Delete("/{id}", bookHandler.DeleteBook)
			r.Post("/payment_methods", paymentHandler.CreatePaymentMethod)
		})
	})

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return cors(r)
}
