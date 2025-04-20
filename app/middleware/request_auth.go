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

func authenticateUser(c *fiber.Ctx, authHeader string, cfg *config.Config) error {
	// 1. 헤더 포맷 검증
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Authorization header format")
	}

	// 2. 토큰 검증
	tokenString := parts[1]
	token, err := validateToken(tokenString, cfg)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Token validation failed: "+err.Error())
	}

	// 3. 클레임 추출
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims structure")
	}

	// 4. 사용자 ID 추출
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user_id in token")
	}

	// 5. DB에서 사용자 조회
	db, ok := c.Locals("db").(*gorm.DB)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "Database connection not available")
	}

	var user entity.User
	if err := db.First(&user, int(userID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusUnauthorized, "User not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Database query error")
	}

	// 6. 컨텍스트에 사용자 정보 저장
	c.Locals("is_authenticated", true)
	c.Locals("request_user", user)
	return nil
}

func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("is_authenticated", false)
		c.Locals("request_user", "anonymous")

		// Authorization 헤더 확인
		authHeader := c.Get("Authorization", "")

		if authHeader == "" {
			return c.Next() // 인증 없이 계속 진행
		}

		if err := authenticateUser(c, authHeader, cfg); err != nil {
			switch e := err.(type) {
			case *fiber.Error:
				return c.Status(e.Code).JSON(common.CoreResponse{
					Message: err.Error(),
				})
			default:
				c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
					Message: "internal server error",
				})
			}
		}

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
