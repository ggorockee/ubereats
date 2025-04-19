package middleware

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"

	"github.com/gofiber/fiber/v2"
)

func RoleGuard(roles ...entity.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("request_user").(entity.User)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(common.CoreResponse{
				Message: "user not authenticated",
			})
		}

		for _, role := range roles {
			if role == entity.RoleAny {
				return c.Next()
			}
		}

		// 사용자 역할과 허용된 역할 비교
		userRole := user.Role
		for _, allowedRole := range roles {
			if userRole == allowedRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(common.CoreResponse{
			Message: "insufficient role permissions",
		})
	}
}
