package router

import (
	"book-store/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"net/http"
)

func NewRouter(authHandler *handler.AuthHandler,bookHandler *handler.BookHandler, paymentHandler *handler.PaymentMethodHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	r.Group(func(r chi.Router) {
		r.Get("/login", authHandler.Login)
		r.Post("/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
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
