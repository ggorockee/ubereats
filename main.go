package main

import (
	"ubereats/app"
	"ubereats/app/core/helper"
	"ubereats/app/domain/category"
	"ubereats/app/domain/restaurant"
	"ubereats/app/domain/user"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		helper.Module,

		restaurant.ControllerMoudle,
		category.ControllerMoudle,
		user.ControllerMoudle,

		fx.Provide(
			app.NewFiber,
			fx.Annotate(
				app.NewRouter,
				fx.ParamTags(``, `group:"routes"`),
			),
		),

		fx.Invoke(
			func(fiber.Router) {},
			func(*fiber.App) {},
		),
	).Run()
}
