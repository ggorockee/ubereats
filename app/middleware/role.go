package middleware

import (
	"fmt"
	"log"
	"slices"
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

		// 0. 관리자 처리
		if user.Role == entity.RoleAdmin {
			return c.Next()
		}

		// 1. RoleAny 처리
		if slices.Contains(roles, entity.RoleAny) {
			return c.Next()
		}

		// 2. 역할 일치 확인
		if slices.Contains(roles, user.Role) {
			log.Printf("✅ 허용 역할: %v, 사용자 역할: %s", roles, user.Role)
			return c.Next()
		}

		// 3. 권한 없음 (핵심 수정 부분)
		log.Printf("⛔ 거부됨 - 허용 역할: %v, 사용자 역할: %s", roles, user.Role)
		return c.Status(fiber.StatusForbidden).JSON(common.CoreResponse{
			Message: fmt.Sprintf(
				"허용 역할: %v, 현재 역할: %s",
				roles,
				user.Role,
			),
		})

	}

}
