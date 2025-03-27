package controller

import (
	"ubereats/app"

	"ubereats/app/core/helper/common"
	authDto "ubereats/app/domain/auth/dto"
	authService "ubereats/app/domain/auth/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Table() []app.Mapping
	SignUp(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authController struct {
	authService authService.AuthService
}

// Login
// @Summary 로그인
// @Description 로그인
// @Tags Auth
// @Accept json
// @Produce json
// @Param requestBody body authDto.LoginInput true "requestBody"
// @Router /auth/login [post]
// @Security Bearer
func (ctrl *authController) Login(c *fiber.Ctx) error {
	var requestBody authDto.LoginInput
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common.CoreResponse{
				Message: err.Error(),
			},
		)
	}

	if err := common.ValidateStruct(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			common.CoreResponse{
				Message: err.Error(),
			},
		)
	}

	output, err := ctrl.authService.Login(c, &requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// SignUp
// @Summary 계정 생성
// @Description 계정 생성
// @Tags Auth
// @Accept json
// @Produce json
// @Param requestBody body authDto.SignUpInput true "requestBody"
// @Router /auth/signup [post]
// @Security Bearer
func (ctrl *authController) SignUp(c *fiber.Ctx) error {

	var requestBody authDto.SignUpInput
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common.CoreResponse{
				Message: err.Error(),
			},
		)
	}

	if err := common.ValidateStruct(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			common.CoreResponse{
				Message: err.Error(),
			},
		)
	}

	output, err := ctrl.authService.SignUp(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// Table implements AuthController.
func (ctrl *authController) Table() []app.Mapping {
	v1 := "/api/v1/auth"
	return []app.Mapping{
		{
			Path:    v1 + "/signup",
			Method:  fiber.MethodPost,
			Handler: ctrl.SignUp,
		},
		{
			Path:    v1 + "/login",
			Method:  fiber.MethodPost,
			Handler: ctrl.Login,
		},
	}
}

func NewAuthController(
	authService authService.AuthService,
) AuthController {
	return &authController{
		authService: authService,
	}
}
