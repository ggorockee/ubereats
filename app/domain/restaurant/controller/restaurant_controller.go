package controller

import (
	"log"
	"ubereats/app"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantSvc "ubereats/app/domain/restaurant/service"
	"ubereats/app/middleware"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
)

type RestaurantController interface {
	Table() []app.Mapping
	CreateRestaurant(c *fiber.Ctx) error
	EditRestaurant(c *fiber.Ctx) error
	GetAllRestaurant(c *fiber.Ctx) error
}

type restaurantController struct {
	restaurantService restaurantSvc.RestaurantService
	cfg               *config.Config
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
		log.Println("너니?")
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

// GetAllRestaurant
// @Summary GetAllRestaurant
// @Description GetAllRestaurant
// @Tags Restaurant
// @Accept json
// @Produce json
// @Router /restaurant [get]
// @Security Bearer
func (ctrl *restaurantController) GetAllRestaurant(c *fiber.Ctx) error {
	output, err := ctrl.restaurantService.GetAllRestaurant(c)
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
			Handler: ctrl.GetAllRestaurant,
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
