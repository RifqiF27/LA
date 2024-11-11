package router

import (
	"book-store/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"net/http"
)

func NewRouter(authHandler *handler.AuthHandler, paymentHandler *handler.PaymentMethodHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Group(func(r chi.Router) {
		r.Get("/login", authHandler.Login)
		r.Post("/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Get("/dashboard", authHandler.Dashboard)
		r.Post("/payment_methods", paymentHandler.CreatePaymentMethod)
	})

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return cors(r)
}
