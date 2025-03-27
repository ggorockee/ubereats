package app

import (
	"context"
	"fmt"
	"log"
	"time"
	"ubereats/app/middleware"
	"ubereats/config"

	_ "ubereats/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewFiber(lc fx.Lifecycle, cfg *config.Config, db *gorm.DB) *fiber.App {
	app := initializeFiber(cfg, db)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := fmt.Sprintf(":%s", cfg.Server.Port)
			go func() {
				if err := app.Listen(addr); err != nil {
					log.Println("fiber server error", zap.Error(err))
				}
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
				log.Println("Fiber server started successfully")
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down Fiber server")
			return app.Shutdown()
		},
	})

	return app
}

func initializeFiber(cfg *config.Config, db *gorm.DB) *fiber.App {
	app := fiber.New()
	app.Get("/api/v1/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	app.Get("/api/v1/docs/*", swagger.HandlerDefault)

	app = setMiddleware(app, cfg, db)

	return app
}

func setMiddleware(app *fiber.App, cfg *config.Config, db *gorm.DB) *fiber.App {
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})
	app.Use(middleware.AuthMiddleware(cfg))
	// app.Use(recover.New())
	return app
}
