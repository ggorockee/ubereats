package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewRouter(app *fiber.App, routes []Route) fiber.Router {
	for _, route := range routes {
		for _, mapping := range route.Table() {
			handlers := append(mapping.Middlewares, mapping.Handler)

			switch mapping.Method {
			case fiber.MethodGet:
				app.Get(mapping.Path, handlers...)
			case fiber.MethodPost:
				app.Post(mapping.Path, handlers...)
			case fiber.MethodPut:
				app.Put(mapping.Path, handlers...)
			case fiber.MethodPatch:
				app.Patch(mapping.Path, handlers...)
			case fiber.MethodDelete:
				app.Delete(mapping.Path, handlers...)
			}

		}

	}
	return app
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

type Route interface {
	Table() []Mapping
}

type Mapping struct {
	Method      string
	Path        string
	Handler     func(ctx *fiber.Ctx) error
	Middlewares []func(*fiber.Ctx) error
}
