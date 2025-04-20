package middleware

import (
	"ubereats/app/core/helper/common"
	"ubereats/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWtProtected(cfg *config.Config) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   jwtware.SigningKey{Key: []byte(cfg.Secret.Jwt)},
		AuthScheme:   "Bearer",
		TokenLookup:  "header:Authorization",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(
			common.CoreResponse{
				Ok:      false,
				Message: err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(
		common.CoreResponse{
			Ok:      false,
			Message: "Invalid or expired JWT",
		},
	)
}
