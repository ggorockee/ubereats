package controller

import (
	"strconv"
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	dishDto "ubereats/app/domain/dish/dto"
	dishSvc "ubereats/app/domain/dish/service"
	"ubereats/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type DishController interface {
	Table() []app.Mapping
	GetAllDish(c *fiber.Ctx) error
	CreateDish(c *fiber.Ctx) error
	UpdateDish(c *fiber.Ctx) error
	DeleteDish(c *fiber.Ctx) error
	FindDishById(c *fiber.Ctx) error
	SearchDish(c *fiber.Ctx) error
	GetAllCategories(c *fiber.Ctx) error
}

type dishController struct {
	dishSvc dishSvc.DishService
	cfg     *config.Config
}

// FindDishById implements DishController.
func (ctrl *dishController) FindDishById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	input := dishDto.DishInput{DishID: id}
	output, err := ctrl.dishSvc.FindDishById(input)
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

// SearchDish implements DishController.
func (ctrl *dishController) SearchDish(c *fiber.Ctx) error {
	query := c.Query("query")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := dishDto.SearchDish{Query: query, Page: page}

	output, err := ctrl.dishSvc.SearchDishByName(input)
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

// GetAllCategories implements DishController.
func (ctrl *dishController) GetAllCategories(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := dishDto.DishsInput{Page: page}
	output, err := ctrl.dishSvc.AllDishs(input)
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

// GetAllDish implements DishController.
func (ctrl *dishController) GetAllDish(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := dishDto.DishsInput{Page: page}
	output, err := ctrl.dishSvc.AllDishs(input)
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

func (ctrl *dishController) DeleteDish(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = ctrl.dishSvc.DeleteDish(id, c)
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

// UpdateDish implements DishController.
func (ctrl *dishController) UpdateDish(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var requestBody dishDto.UpdateDish
	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dish, err := ctrl.dishSvc.UpdateDish(&requestBody, id, c)
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
		Data:    dish,
	})
}

func (ctrl *dishController) CreateDish(c *fiber.Ctx) error {
	var requestBody dishDto.CreateDish
	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dish, err := ctrl.dishSvc.CreateDish(&requestBody, c)
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
		Data:    dish,
	})

}

// Table implements DishController.
func (ctrl *dishController) Table() []app.Mapping {
	v1 := "/api/v1/dish"
	v2 := "/api/v1/category"

	return []app.Mapping{
		{Method: fiber.MethodGet, Path: v1 + "", Handler: ctrl.GetAllDish},
		{Method: fiber.MethodGet, Path: v1 + "/:id", Handler: ctrl.FindDishById},
		{Method: fiber.MethodGet, Path: v1 + "/search", Handler: ctrl.SearchDish},
		{
			Method:  fiber.MethodPost,
			Path:    v1 + "",
			Handler: ctrl.CreateDish,
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
			Handler: ctrl.UpdateDish,
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
			Handler: ctrl.DeleteDish,
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

func NewDishController(r dishSvc.DishService, c *config.Config) DishController {
	return &dishController{
		dishSvc: r,
		cfg:     c,
	}
}
