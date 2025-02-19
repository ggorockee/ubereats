package controller

import (
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/helper/common"
	categoryDto "ubereats/app/domain/category/dto"
	categoryRes "ubereats/app/domain/category/response"
	categorySvc "ubereats/app/domain/category/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Table() []app.Mapping
	GetAllCategory(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	GetFindById(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type categoryController struct {
	categorySvc categorySvc.CategoryService
	cfg         *config.Config
}

// DeleteCategory implements CategoryController.
func (ctrl *categoryController) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = ctrl.categorySvc.DeleteCategory(id, c)
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
		Data:    nil,
	})
}

// GetFindById implements CategoryController.
func (ctrl *categoryController) GetFindById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	category, err := ctrl.categorySvc.GetFindById(id, c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	categoryResponse := categoryRes.GenCategoryRes(category)
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    categoryResponse,
	})
}

// UpdateCategory implements CategoryController.
func (ctrl *categoryController) UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var requestBody categoryDto.UpdateCategory
	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	category, err := ctrl.categorySvc.UpdateCategory(&requestBody, id, c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	categoryResponse := categoryRes.GenCategoryRes(category)
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    categoryResponse,
	})
}

func (ctrl *categoryController) CreateCategory(c *fiber.Ctx) error {
	var requestBody categoryDto.CreateCategory
	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	category, err := ctrl.categorySvc.CreateCategory(&requestBody, c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	categoryResponse := categoryRes.GenCategoryRes(category)
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    categoryResponse,
	})

}

// GetAll implements CategoryController.
func (ctrl *categoryController) GetAllCategory(c *fiber.Ctx) error {
	categories, err := ctrl.categorySvc.GetAllCategory(c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	categoriesResponse := categoryRes.GenCategoriesRes(categories)
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    categoriesResponse,
	})
}

// Table implements CategoryController.
func (ctrl *categoryController) Table() []app.Mapping {
	v1 := "/api/v1/category"

	return []app.Mapping{
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "",
			Handler: ctrl.GetAllCategory,
		},
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "/:id",
			Handler: ctrl.GetFindById,
		},
		{
			Method:  fiber.MethodPost,
			Path:    v1 + "",
			Handler: ctrl.CreateCategory,
			Middlewares: []fiber.Handler{
				app.AuthMiddleware(ctrl.cfg),
			},
		},
		{
			Method:  fiber.MethodPut,
			Path:    v1 + "/:id",
			Handler: ctrl.UpdateCategory,
			Middlewares: []fiber.Handler{
				app.AuthMiddleware(ctrl.cfg),
			},
		},
		{
			Method:  fiber.MethodDelete,
			Path:    v1 + "/:id",
			Handler: ctrl.DeleteCategory,
			Middlewares: []fiber.Handler{
				app.AuthMiddleware(ctrl.cfg),
			},
		},
	}
}

func NewCategoryController(r categorySvc.CategoryService, c *config.Config) CategoryController {
	return &categoryController{
		categorySvc: r,
		cfg:         c,
	}
}
