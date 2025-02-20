package controller

import (
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/entity"
	categorySvc "ubereats/app/domain/category/service"
	"ubereats/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Table() []app.Mapping
	CreateCategory(c *fiber.Ctx) error
	EditCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
	AllCategories(c *fiber.Ctx) error
	FindCategoryById(c *fiber.Ctx) error
	SearchCategoryByName(c *fiber.Ctx) error
}

type categoryController struct {
	categorySvc categorySvc.CategoryService
	cfg         *config.Config
}

// AllCategories implements CategoryController.
func (ctrl *categoryController) AllCategories(c *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateCategory implements CategoryController.
func (ctrl *categoryController) CreateCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryController.
func (ctrl *categoryController) DeleteCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// EditCategory implements CategoryController.
func (ctrl *categoryController) EditCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FindCategoryById implements CategoryController.
func (ctrl *categoryController) FindCategoryById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// SearchCategoryByName implements CategoryController.
func (ctrl *categoryController) SearchCategoryByName(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Table implements CategoryController.
func (ctrl *categoryController) Table() []app.Mapping {
	v1 := "/api/v1/category"

	return []app.Mapping{
		{Method: fiber.MethodGet, Path: v1 + "s", Handler: ctrl.AllCategories},
		{Method: fiber.MethodGet, Path: v1 + "/:id", Handler: ctrl.FindCategoryById},
		{Method: fiber.MethodGet, Path: v1 + "/search", Handler: ctrl.SearchCategoryByName},
		{
			Method:  fiber.MethodPost,
			Path:    v1,
			Handler: ctrl.CreateCategory,
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
			Handler: ctrl.EditCategory,
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
			Handler: ctrl.DeleteCategory,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{
					entity.RoleOwner,
				}),
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
