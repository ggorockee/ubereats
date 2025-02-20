package controller

import (
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/entity"
	dishSvc "ubereats/app/domain/dish/service"
	"ubereats/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type DishController interface {
	Table() []app.Mapping

	AllCategories(c *fiber.Ctx) error
	FindDishById(c *fiber.Ctx) error
	SearchDishByName(c *fiber.Ctx) error
}

type dishController struct {
	dishSvc dishSvc.DishService
	cfg     *config.Config
}

// AllCategories implements DishController.
func (ctrl *dishController) AllCategories(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FindDishById implements DishController.
func (ctrl *dishController) FindDishById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// SearchDishByName implements DishController.
func (ctrl *dishController) SearchDishByName(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Table implements DishController.
func (ctrl *dishController) Table() []app.Mapping {
	v1 := "/api/v1/dish"

	return []app.Mapping{
		{Method: fiber.MethodGet, Path: v1 + "", Handler: ctrl.AllDishes},
		{Method: fiber.MethodGet, Path: v1 + "/:id", Handler: ctrl.FindDishById},
		{Method: fiber.MethodGet, Path: v1 + "/search", Handler: ctrl.SearchDishByName},
		{
			Method:  fiber.MethodPost,
			Path:    v1,
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
			Handler: ctrl.EditDish,
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
	}
}

func NewDishController(r dishSvc.DishService, c *config.Config) DishController {
	return &dishController{
		dishSvc: r,
		cfg:     c,
	}
}
