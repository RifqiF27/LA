//go:build wireinject
// +build wireinject

package wire

import (
	"net/http"
	"os"
	"travelika/database"
	"travelika/handler"
	"travelika/repository"
	"travelika/router"
	"travelika/service"

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
		repository.NewSessionRepository,
		repository.NewDestinationRepository,
		repository.NewTransactionRepository,
		service.NewUserService,
		service.NewDestinationService,
		service.NewTransactionService,
		handler.NewAuthHandler,
		handler.NewDestinationHandler,
		handler.NewTransactionHandler,
		router.NewRouter,
	)
	return nil, nil
}
