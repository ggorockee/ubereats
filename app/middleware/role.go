package middleware

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"

	"github.com/gofiber/fiber/v2"
)

type AllowedRoles []entity.UserRole

func Role(roles AllowedRoles) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 인증된 사용자 가져오기
		user, ok := c.Locals("request_user").(entity.User)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(
				common.CoreResponse{
					Message: "user not authenticated",
				},
			)
		}

		for _, role := range roles {
			if role == entity.RoleAny {
				// 모든 역할 허용
				return c.Next()
			}
		}

		// 사용자 역할과 허용된 역할 비교
		userRole := user.Role
		for _, allowedRole := range roles {
			if userRole == allowedRole {
				return c.Next() // 역할 일치시 다음 핸들러
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(
			common.CoreResponse{
				Message: "insufficient role permissions",
			},
		)
	}
}
