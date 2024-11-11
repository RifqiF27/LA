//go:build wireinject
// +build wireinject


package wire

import (
	"book-store/database"
	"book-store/handler"
	"book-store/repository"
	"book-store/router"
	"book-store/service"
	"net/http"

	// "github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func InitializeRouterHandler() http.Handler {
	wire.Build(
		database.NewPostgresDB,
		repository.NewUserRepo,
		repository.NewPaymentMethodRepository,
		service.NewUserService,
		service.NewPaymentMethodService,
		handler.NewAuthHandler,
		handler.NewPaymentMethodHandler,
		router.NewRouter,
	)
	return nil
}
