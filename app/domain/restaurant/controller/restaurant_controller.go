package controller

import (
	"log"
	"ubereats/app"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantSvc "ubereats/app/domain/restaurant/service"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
)

type RestaurantController interface {
	Table() []app.Mapping
	CreateRestaurant(c *fiber.Ctx) error
}

type restaurantController struct {
	restaurantService restaurantSvc.RestaurantService
	cfg               *config.Config
}

func (ctrl *restaurantController) CreateRestaurant(c *fiber.Ctx) error {
	var inputParam restaurantDto.CreateRestaurantDto
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
	log.Println(v1)

	return []app.Mapping{
		{
			Method: fiber.MethodGet,
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
