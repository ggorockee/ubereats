package controller

import (
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRes "ubereats/app/domain/restaurant/response"
	restaurantSvc "ubereats/app/domain/restaurant/service"
	"ubereats/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type RestaurantController interface {
	Table() []app.Mapping
	GetAllRestaurnat(c *fiber.Ctx) error
	CreateRestaurant(c *fiber.Ctx) error
	UpdateRestaurant(c *fiber.Ctx) error
}

type restaurantController struct {
	restaurantSvc restaurantSvc.RestaurantService
	cfg           *config.Config
}

// UpdateRestaurant implements RestaurantController.
func (ctrl *restaurantController) UpdateRestaurant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var requestBody restaurantDto.UpdateRestaurant
	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	restaurant, err := ctrl.restaurantSvc.UpdateRestaurant(&requestBody, id, c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	restaurantResponse := restaurantRes.GenRestaurantRes(restaurant)
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    restaurantResponse,
	})
}

func (ctrl *restaurantController) CreateRestaurant(c *fiber.Ctx) error {
	var requestBody restaurantDto.CreateRestaurant
	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	restaurant, err := ctrl.restaurantSvc.CreateRestaurant(&requestBody, c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	restaurantResponse := restaurantRes.GenRestaurantRes(restaurant)
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    restaurantResponse,
	})

}

// GetAll implements RestaurantController.
func (ctrl *restaurantController) GetAllRestaurnat(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Table implements RestaurantController.
func (ctrl *restaurantController) Table() []app.Mapping {
	v1 := "/api/v1/restaurant"

	return []app.Mapping{
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "",
			Handler: ctrl.GetAllRestaurnat,
		},

		{
			Method:  fiber.MethodPost,
			Path:    v1 + "",
			Handler: ctrl.CreateRestaurant,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{
					entity.RoleOwner,
				}),
			},
		},

		{
			Method:  fiber.MethodPut,
			Path:    v1 + "/:id",
			Handler: ctrl.UpdateRestaurant,
		},
	}
}

func NewRestaurantController(r restaurantSvc.RestaurantService, c *config.Config) RestaurantController {
	return &restaurantController{
		restaurantSvc: r,
		cfg:           c,
	}
}
