package middleware

import (
	"errors"
	"strings"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("is_authenticated", false)
		c.Locals("request_user", "anonymous")

		// Authorization 헤더 확인
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return common.ErrorResponse(c, common.ErrArg{
				IsError: true,
				Code:    fiber.StatusUnauthorized,
				Message: "Authorization header required",
				Data:    nil,
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {

			return common.ErrorResponse(c, common.ErrArg{
				IsError: true,
				Code:    fiber.StatusUnauthorized,
				Message: "Invalid Authorization header format",
				Data:    nil,
			})
		}

		tokenString := parts[1]

		// 토큰 검증
		token, err := validateToken(tokenString, cfg)
		if err != nil {

			return common.ErrorResponse(c, common.ErrArg{
				IsError: true,
				Code:    fiber.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})

		}

		// Claims 처리
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {

			return common.ErrorResponse(c, common.ErrArg{
				IsError: true,
				Code:    fiber.StatusUnauthorized,
				Message: "Invalid token claims",
				Data:    nil,
			})
		}

		userIdFloat, ok := claims["user_id"].(float64)
		if !ok {

			return common.ErrorResponse(c, common.ErrArg{
				IsError: true,
				Code:    fiber.StatusUnauthorized,
				Message: "Invalid user_id in token",
				Data:    nil,
			})
		}
		userId := int(userIdFloat)

		// DB 조회
		db, ok := c.Locals("db").(*gorm.DB)
		if !ok {

			return common.ErrorResponse(c, common.ErrArg{
				IsError: true,
				Code:    fiber.StatusInternalServerError,
				Message: "Database connection not available",
				Data:    nil,
			})
		}

		var user entity.User
		if err := db.Where("id = ?", userId).First(&user).Error; err != nil {

			return common.ErrorResponse(c, common.ErrArg{
				IsError: true,
				Code:    fiber.StatusUnauthorized,
				Message: "User not found",
				Data:    nil,
			})
		}

		c.Locals("is_authenticated", true)
		c.Locals("request_user", user)

		return c.Next()

	}
}

func validateToken(tokenString string, cfg *config.Config) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(cfg.Secret.Jwt), nil
	})

	if err != nil {
		switch err {
		case jwt.ErrSignatureInvalid:
			return nil, errors.New("invalid token signature")
		case jwt.ErrTokenExpired:
			return nil, errors.New("token has expired")
		default:
			return nil, errors.New("token validation failed")
		}
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return token, nil
}
