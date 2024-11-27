package main

import (
	"ecommerce/infra"
	"ecommerce/router"

	"go.uber.org/zap"
)

func main() {
	ctx, err := infra.NewContext()
	if err != nil {
		ctx.Log.Panic("Error", zap.Error(err))
		return
	}

	router.SetupReouter(ctx)
}
