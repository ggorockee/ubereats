package main

import (
	"ubereats/app"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,

		fx.Provide(
			app.NewFiber,
		),

		fx.Invoke(
			func(*fiber.App) {},
		),
	).Run()
}
