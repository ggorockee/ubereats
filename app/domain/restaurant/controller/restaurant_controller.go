package controller

import (
	"strconv"
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantSvc "ubereats/app/domain/restaurant/service"
	"ubereats/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type RestaurantController interface {
	Table() []app.Mapping
	GetAllRestaurant(c *fiber.Ctx) error
	CreateRestaurant(c *fiber.Ctx) error
	UpdateRestaurant(c *fiber.Ctx) error
	DeleteRestaurant(c *fiber.Ctx) error
	GetFindById(c *fiber.Ctx) error
	SearchRestaurant(c *fiber.Ctx) error
	GetAllCategories(c *fiber.Ctx) error
}

type restaurantController struct {
	restaurantSvc restaurantSvc.RestaurantService
	cfg           *config.Config
}

// GetAllCategories implements RestaurantController.
func (ctrl *restaurantController) GetAllCategories(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := restaurantDto.RestaurantsInput{Page: page}
	output, err := ctrl.restaurantSvc.AllRestaurants(input)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    output,
	})
}

// GetAllRestaurant implements RestaurantController.
func (ctrl *restaurantController) GetAllRestaurant(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := restaurantDto.RestaurantsInput{Page: page}
	output, err := ctrl.restaurantSvc.AllRestaurants(input)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    output,
	})
}

// SearchRestaurant implements RestaurantController.
func (ctrl *restaurantController) SearchRestaurant(c *fiber.Ctx) error {
	query := c.Query("query")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := restaurantDto.SearchRestaurant{Query: query, Page: page}

	output, err := ctrl.restaurantSvc.SearchRestaurantByName(input)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    output,
	})
}

func (ctrl *restaurantController) GetFindById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	restaurant, err := ctrl.restaurantSvc.GetFindById(id, c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    restaurant,
	})
}

func (ctrl *restaurantController) DeleteRestaurant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = ctrl.restaurantSvc.DeleteRestaurant(id, c)
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    nil,
	})

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

	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    restaurant,
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

	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    restaurant,
	})

}

// GetAll implements RestaurantController.
func (ctrl *restaurantController) GetAllRestaurnat(c *fiber.Ctx) error {
	restaurants, err := ctrl.restaurantSvc.GetAllRestaurant(c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    restaurants,
	})

}

// Table implements RestaurantController.
func (ctrl *restaurantController) Table() []app.Mapping {
	v1 := "/api/v1/restaurant"
	v2 := "/api/v1/category"

	return []app.Mapping{
		{Method: fiber.MethodGet, Path: v1 + "", Handler: ctrl.GetAllRestaurnat},
		{Method: fiber.MethodGet, Path: v1 + "/:id", Handler: ctrl.GetFindById},
		{Method: fiber.MethodGet, Path: v1 + "/search", Handler: ctrl.SearchRestaurant},
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
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{
					entity.RoleOwner,
				}),
			},
		},
		{
			Method:  fiber.MethodDelete,
			Path:    v1 + "/:id",
			Handler: ctrl.DeleteRestaurant,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{
					entity.RoleOwner,
				}),
			},
		},
		{Method: fiber.MethodGet, Path: v2, Handler: ctrl.GetAllCategories},
	}
}

func NewRestaurantController(r restaurantSvc.RestaurantService, c *config.Config) RestaurantController {
	return &restaurantController{
		restaurantSvc: r,
		cfg:           c,
	}
}
