package controller

import (
	"ubereats/app"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantResp "ubereats/app/domain/restaurant/response"
	restaurantSvc "ubereats/app/domain/restaurant/service"
	"ubereats/app/middleware"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
)

type RestaurantController interface {
	Table() []app.Mapping
	CreateRestaurant(c *fiber.Ctx) error
	EditRestaurant(c *fiber.Ctx) error
	AllRestaurant(c *fiber.Ctx) error
	GetCategoryByName(c *fiber.Ctx) error
}

type restaurantController struct {
	restaurantService restaurantSvc.RestaurantService
	cfg               *config.Config
}

// GetCategoryByName
// @Summary Get category by name
// @Description Get restaurants in a category with pagination
// @Tags Restaurant
// @Accept json
// @Produce json
// @Param name query string true "Category name"
// @Param page query int false "Page number (default: 1)"
// @Router /api/v1/restaurant/category [get]
// @Security Beare
func (ctrl *restaurantController) GetCategoryByName(c *fiber.Ctx) error {
	var input restaurantDto.GetCategoryIn
	// 1. 쿼리 파라미터 파싱
	if err := c.QueryParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(restaurantResp.GetCategoryOut{
			Ok:      false,
			Message: "Invalid query parameters",
		})
	}

	// 2. 필수 필드 검증 (name 필수)
	if err := common.ValidateStruct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(restaurantResp.GetCategoryOut{
			Ok:      false,
			Message: err.Error(),
		})
	}

	// 3. 페이지네이션 기본값 설정
	params := common.PaginationParams{
		Page:  input.Page,
		Limit: input.Limit, // Limit이 0이면 서비스에서 기본값 처리
	}

	// 4. 서비스 호출
	output, err := ctrl.restaurantService.FindCategoryByName(c, input.Name, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// EditRestaurant
// @Summary EditRestaurant
// @Description EditRestaurant
// @Tags Restaurant
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param requestBody body dto.EditRestaurantIn true "requestBody"
// @Router /restaurant/{id} [put]
// @Security Bearer
func (ctrl *restaurantController) EditRestaurant(c *fiber.Ctx) error {
	var requestBody restaurantDto.EditRestaurantIn
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

	restaurantId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	output, err := ctrl.restaurantService.EditRestaurant(c, uint(restaurantId), &requestBody)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// AllRestaurant
// @Summary AllRestaurant
// @Description AllRestaurant
// @Tags Restaurant
// @Accept json
// @Produce json
// @Router /restaurant [get]
// @Security Bearer
func (ctrl *restaurantController) AllRestaurant(c *fiber.Ctx) error {
	var input restaurantDto.GetRestaurantsInput
	if err := c.QueryParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(restaurantResp.AllRestaurantOut{
			Message: "Invalid query parameters",
		})
	}

	output, err := ctrl.restaurantService.AllRestaurant(c, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// CreateRestaurant
// @Summary Restaurant
// @Description Restaurant
// @Tags Restaurant
// @Accept json
// @Produce json
// @Param requestBody body dto.CreateRestaurantIn true "requestBody"
// @Router /restaurant [post]
// @Security Bearer
func (ctrl *restaurantController) CreateRestaurant(c *fiber.Ctx) error {
	var inputParam restaurantDto.CreateRestaurantIn

	if err := c.BodyParser(&inputParam); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	if err := common.ValidateStruct(&inputParam); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.CoreResponse{
			Message: err.Error(),
		})
	}

	output, err := ctrl.restaurantService.CreateRestaurant(c, &inputParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (ctrl *restaurantController) Table() []app.Mapping {
	v1 := "/api/v1/restaurant"

	return []app.Mapping{
		{
			Method:  fiber.MethodPost,
			Path:    v1 + "",
			Handler: ctrl.CreateRestaurant,
			Middlewares: []fiber.Handler{
				middleware.JWtProtected(ctrl.cfg),
				middleware.RoleGuard(entity.RoleOwner),
			},
		},
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "",
			Handler: ctrl.AllRestaurant,
			// Middlewares: []fiber.Handler{
			// 	// middleware.RoleGuard(entity.RoleOwner),
			// },
		},
		{
			Method:  fiber.MethodPut,
			Path:    v1 + "/:id",
			Handler: ctrl.EditRestaurant,
			Middlewares: []fiber.Handler{
				middleware.JWtProtected(ctrl.cfg),
				middleware.RoleGuard(entity.RoleOwner),
			},
		},
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "/category",
			Handler: ctrl.GetCategoryByName,
			Middlewares: []fiber.Handler{
				middleware.JWtProtected(ctrl.cfg),
				middleware.RoleGuard(entity.RoleAny),
			},
		},
	}
}

func NewRestaurantController(
	restaurantService restaurantSvc.RestaurantService,
	cfg *config.Config,
) RestaurantController {
	return &restaurantController{
		cfg:               cfg,
		restaurantService: restaurantService,
	}
}
