package controller

import (
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantService "ubereats/app/domain/restaurant/service"
	"ubereats/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Table() []app.Mapping
	CreateCategory(c *fiber.Ctx) error
}

type categoryController struct {
	cfg             *config.Config
	categoryService restaurantService.CategoryService
}

// Table implements CategoryController.
func (ctrl *categoryController) Table() []app.Mapping {
	v1 := "/api/v1/category"
	return []app.Mapping{
		{
			Path:    v1 + "",
			Handler: ctrl.CreateCategory,
			Method:  fiber.MethodPost,
			Middlewares: []fiber.Handler{
				middleware.Role(middleware.AllowedRoles{
					entity.RoleOwner,
				}),
			},
		},
	}
}

func NewCategoryController(
	cfg *config.Config,
	categoryService restaurantService.CategoryService,

) CategoryController {
	return &categoryController{
		cfg:             cfg,
		categoryService: categoryService,
	}
}

// CreateCategory
// @Summary 로그인
// @Description 로그인
// @Tags Category
// @Accept json
// @Produce json
// @Param requestBody body restaurantDto.CreateCategoryInput true "requestBody"
// @Router /auth/login [post]
// @Security Bearer
func (ctrl *categoryController) CreateCategory(c *fiber.Ctx) error {
	var requestBody restaurantDto.CreateCategoryInput
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common.CoreResponse{
				Message: err.Error(),
			},
		)
	}

	output, err := ctrl.categoryService.CreateCategory(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
