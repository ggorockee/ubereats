package main

import (
	"ubereats/app"
	"ubereats/app/core/helper"
	"ubereats/app/domain/user"

	"ubereats/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// @title ggorockee App
// @version 1.0
// @description This is an API for Truloop Application
// @contact.name ggorockee
// @contact.email ggorockee@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	fx.New(
		config.Module,
		helper.Module,

		user.ControllerModule,
		// restaurant.ControllerModule,

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
