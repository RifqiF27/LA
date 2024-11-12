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
	"os"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ProvideLogger() (*zap.Logger, error) {

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(file),
			zap.InfoLevel,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zap.InfoLevel,
		),
	)

	logger := zap.New(core)
	return logger, nil
}

func InitializeRouterHandler() (http.Handler, error) {
	wire.Build(
		database.NewPostgresDB,
		ProvideLogger,
		repository.NewUserRepo,
		repository.NewPaymentMethodRepository,
		repository.NewBookRepository, 
		service.NewUserService,
		service.NewPaymentMethodService,
		service.NewBookService, 
		handler.NewAuthHandler,
		handler.NewPaymentMethodHandler,
		handler.NewBookHandler, 
		router.NewRouter,
	)
	return nil, nil
}
