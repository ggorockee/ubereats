package controller

import (
	"ubereats/app"
	"ubereats/app/core/helper/common"
	userSvc "ubereats/app/domain/user/service"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Table() []app.Mapping
	Hi(c *fiber.Ctx) error
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

func (ctrl *userController) Table() []app.Mapping {
	v1 := "/api/v1/hi"
	return []app.Mapping{
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "",
			Handler: ctrl.Hi,
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
