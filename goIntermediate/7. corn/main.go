package main

import (
	"log"
	"voucher_system/infra"
	"voucher_system/router"

	_ "voucher_system/docs"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// @title Voucher System API
// @version 1.0
// @description API for managing vouchers
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Authentication
// @in header
// @name Authorization
// @securityDefinitions.apikey UserID
// @in header
// @name User-ID

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := router.NewRoutes(*ctx)

	go startCronJob(ctx)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func startCronJob(ctx *infra.ServiceContext) {
	c := cron.New()

	_, err := c.AddFunc("@every 20s", func() {
		ctx.Crn.RunVoucherUsageJob()
	})
	if err != nil {
		ctx.Log.Fatal("Failed to schedule voucher usage job", zap.Error(err))
	}

	ctx.Log.Info("Voucher usage cron job scheduled")
	c.Start()
}
