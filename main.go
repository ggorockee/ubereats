package main

import (
	"ubereats/app"
	"ubereats/app/core/helper"
	"ubereats/app/domain/restaurant"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		helper.Module,

		restaurant.ControllerMoudle,

		fx.Provide(
			app.NewFiber,
		),

		fx.Invoke(
			func(*fiber.App) {},
		),
	).Run()
}
