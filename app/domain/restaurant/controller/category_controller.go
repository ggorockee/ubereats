package controller

import (
	"ubereats/app"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantSvc "ubereats/app/domain/restaurant/service"
	"ubereats/app/middleware"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
)

type CategoryContoller interface {
	Table() []app.Mapping
	CreateCategory(c *fiber.Ctx) error
	GetAllCategory(c *fiber.Ctx) error
	GetCategory(c *fiber.Ctx) error
}

type categoryController struct {
	categoryService restaurantSvc.CategoryService
	cfg             *config.Config
}

// CreateCategory
// @Summary CreateCategory
// @Description CreateCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param requestBody body dto.CreateCategoryIn true "requestBody"
// @Router /category [post]
// @Security Bearer
func (ctrl *categoryController) CreateCategory(c *fiber.Ctx) error {
	var requestBody restaurantDto.CreateCategoryIn

	// json parsing
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	// json validating
	if err := common.ValidateStruct(&requestBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	// service
	output, err := ctrl.categoryService.CreateCategory(c, &requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// GetAllCategory
// @Summary Category
// @Description Category
// @Tags Category
// @Accept json
// @Produce json
// @Router /category [get]
// @Security Bearer
func (ctrl *categoryController) GetAllCategory(c *fiber.Ctx) error {
	output, err := ctrl.categoryService.GetAllCategory(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// GetCategory
// @Summary Category
// @Description Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "category_id"
// @Router /category/{id} [get]
// @Security Bearer
func (ctrl *categoryController) GetCategory(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	output, err := ctrl.categoryService.GetCategory(c, uint(categoryId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// Table implements CategoryContoller.
func (ctrl *categoryController) Table() []app.Mapping {
	v1 := "/api/v1/category"

	return []app.Mapping{
		{
			Path:    v1 + "",
			Method:  fiber.MethodPost,
			Handler: ctrl.CreateCategory,
			Middlewares: []fiber.Handler{
				middleware.JWtProtected(ctrl.cfg),
				middleware.RoleGuard(entity.RoleOwner),
			},
		},
		{
			Path:    v1 + "",
			Method:  fiber.MethodGet,
			Handler: ctrl.GetAllCategory,
		},
		{
			Path:    v1 + "/:id",
			Method:  fiber.MethodGet,
			Handler: ctrl.GetCategory,
		},
	}
}

func NewCategoryContoller(
	categoryService restaurantSvc.CategoryService,
	cfg *config.Config,
) CategoryContoller {
	return &categoryController{
		categoryService: categoryService,
		cfg:             cfg,
	}
}
