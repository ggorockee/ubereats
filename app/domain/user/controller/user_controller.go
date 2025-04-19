package controller

import (
	"ubereats/app"
	"ubereats/app/core/helper/common"
	userDto "ubereats/app/domain/user/dto"
	userSvc "ubereats/app/domain/user/service"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Table() []app.Mapping
	Hi(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
}

type userController struct {
	userService userSvc.UserService
	cfg         *config.Config
}

// Hi
// @Summary Hi
// @Description Hi
// @Tags Dummy
// @Accept json
// @Produce json
// @Router /hi [get]
// @Security Bearer
func (ctrl *userController) Hi(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(common.CoreResponse{
		Ok:      true,
		Message: "hi",
	})
}

// CreateAccount
// @Summary Auth
// @Description Auth
// @Tags Auth
// @Accept json
// @Produce json
// @Param requestBody body dto.CreateAccountIn true "requestBody"
// @Router /user/create [post]
// @Security Bearer
func (ctrl *userController) CreateAccount(c *fiber.Ctx) error {
	var requestBody userDto.CreateAccountIn
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	if err := common.ValidateStruct(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	output, err := ctrl.userService.CreateAccount(c, &requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)

}

func (ctrl *userController) Table() []app.Mapping {
	v1 := "/api/v1"
	return []app.Mapping{
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "/hi",
			Handler: ctrl.Hi,
		},
		{
			Method:  fiber.MethodPost,
			Path:    v1 + "/user/create",
			Handler: ctrl.CreateAccount,
		},
	}
}

func NewUserController(
	userService userSvc.UserService,
	cfg *config.Config,
) UserController {
	return &userController{
		userService: userService,
		cfg:         cfg,
	}
}
