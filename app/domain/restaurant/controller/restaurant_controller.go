package controller

import (
	"ubereats/app"
	"ubereats/config"

	"ubereats/app/core/entity"
	restaurantSvc "ubereats/app/domain/restaurant/service"
	"ubereats/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type RestaurantController interface {
	Table() []app.Mapping
	AllRestaurants(c *fiber.Ctx) error
	FindRestaurantById(c *fiber.Ctx) error
	SearchRestaurantByName(c *fiber.Ctx) error
	CreateRestaurant(c *fiber.Ctx) error
	CreateDish(c *fiber.Ctx) error
	EditRestaurant(c *fiber.Ctx) error
	EditDish(c *fiber.Ctx) error
	DeleteRestaurant(c *fiber.Ctx) error
	DeleteDish(c *fiber.Ctx) error
}

type restaurantController struct {
	restaurantSvc restaurantSvc.RestaurantService
	cfg           *config.Config
}

// AllRestaurants implements RestaurantController.
func (ctrl *restaurantController) AllRestaurants(c *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateDish implements RestaurantController.
func (ctrl *restaurantController) CreateDish(c *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateRestaurant implements RestaurantController.
func (ctrl *restaurantController) CreateRestaurant(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteDish implements RestaurantController.
func (ctrl *restaurantController) DeleteDish(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteRestaurant implements RestaurantController.
func (ctrl *restaurantController) DeleteRestaurant(c *fiber.Ctx) error {
	panic("unimplemented")
}

// EditDish implements RestaurantController.
func (ctrl *restaurantController) EditDish(c *fiber.Ctx) error {
	panic("unimplemented")
}

// EditRestaurant implements RestaurantController.
func (ctrl *restaurantController) EditRestaurant(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FindRestaurantById implements RestaurantController.
func (ctrl *restaurantController) FindRestaurantById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// SearchRestaurantByName implements RestaurantController.
func (ctrl *restaurantController) SearchRestaurantByName(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Table implements RestaurantController.
func (ctrl *restaurantController) Table() []app.Mapping {
	v1 := "/api/v1/restaurant"

	return []app.Mapping{
		{Method: "GET", Path: v1 + "s", Handler: ctrl.AllRestaurants},
		{Method: "GET", Path: v1 + "/:id", Handler: ctrl.FindRestaurantById},
		{Method: "GET", Path: v1 + "/search", Handler: ctrl.SearchRestaurantByName},

		{
			Method:  "POST",
			Path:    v1,
			Handler: ctrl.CreateRestaurant,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},
		{
			Method:  "POST",
			Path:    v1 + "/:id/dish",
			Handler: ctrl.CreateDish,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},

		{
			Method:  "PUT",
			Path:    v1 + "/:id",
			Handler: ctrl.EditRestaurant,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},
		{
			Method:  "PUT",
			Path:    v1 + "/dish/:dishId",
			Handler: ctrl.EditDish,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},

		{
			Method:  "DELETE",
			Path:    v1 + "/:id",
			Handler: ctrl.DeleteRestaurant,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},
		{
			Method:  "DELETE",
			Path:    v1 + "/dish/:dishId",
			Handler: ctrl.DeleteDish,
			Middlewares: []fiber.Handler{
				middleware.AuthMiddleware(ctrl.cfg),
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},
	}
}

func NewRestaurantController(r restaurantSvc.RestaurantService, c *config.Config) RestaurantController {
	return &restaurantController{
		restaurantSvc: r,
		cfg:           c,
	}
}
