package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"ubereats/app"
	"ubereats/app/core/entity"
	"ubereats/app/middleware"
	"ubereats/config"

	"ubereats/app/core/helper/response"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRes "ubereats/app/domain/restaurant/response"
	restaurantSvc "ubereats/app/domain/restaurant/service"
)

type RestaurantController interface {
	Table() []app.Mapping
	CreateRestaurant(c *fiber.Ctx) error
	EditRestaurant(c *fiber.Ctx) error
	DeleteRestaurant(c *fiber.Ctx) error
	SearchRestaurantByName(c *fiber.Ctx) error
	FindRestaurantById(c *fiber.Ctx) error
	AllRestaurants(c *fiber.Ctx) error
}

type restaurantController struct {
	restaurantSvc restaurantSvc.RestaurantService
	cfg           *config.Config
}

func (ctrl *restaurantController) SearchRestaurantByName(c *fiber.Ctx) error {
	query := c.Query("query")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := restaurantDto.SearchRestaurantInput{Query: query, Page: page}
	output, err := ctrl.restaurantSvc.SearchRestaurantByName(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(output)
}

func (ctrl *restaurantController) FindRestaurantById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{Message: "id parsing error"})
	}

	input := restaurantDto.RestaurantInput{RestaurantID: id}
	output, err := ctrl.restaurantSvc.FindRestaurantById(&input)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.BaseResponse{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(output)
}

func (ctrl *restaurantController) AllRestaurants(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	input := restaurantDto.RestaurantsInput{Page: page}
	output, err := ctrl.restaurantSvc.AllRestaurants(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(output)
}

func (ctrl *restaurantController) CreateRestaurant(c *fiber.Ctx) error {
	var requestBody restaurantDto.CreateRestaurantInput
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Message: err.Error(),
		})
	}

	output, err := ctrl.restaurantSvc.CreateRestaurant(&requestBody, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(restaurantRes.CreateRestaurantOutput{
			BaseResponse: response.BaseResponse{Message: err.Error()},
		})
	}

	//return c.Status(fiber.StatusOK).JSON(restaurantRes.CreateRestaurantOutput{
	//	BaseResponse: response.BaseResponse{Ok: true, Data: output},
	//})
	return c.Status(fiber.StatusOK).JSON(output)
}

func (ctrl *restaurantController) EditRestaurant(c *fiber.Ctx) error {
	var input restaurantDto.EditRestaurantInput
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{Message: "id parsing error"})
	}

	input.RestaurantID = id
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{Message: err.Error()})
	}

	output, err := ctrl.restaurantSvc.EditRestaurant(&input, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func (ctrl *restaurantController) DeleteRestaurant(c *fiber.Ctx) error {
	var input restaurantDto.DeleteRestaurantInput
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{Message: "id parsing error"})
	}

	input.RestaurantID = id
	output, err := ctrl.restaurantSvc.DeleteRestaurant(&input, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(output)

}

// AllRestaurants implements RestaurantController.

// Table implements RestaurantController.
func (ctrl *restaurantController) Table() []app.Mapping {
	v1 := "/api/v1/restaurant"
	log.Println(v1)

	return []app.Mapping{
		{Method: fiber.MethodGet, Path: v1 + "", Handler: ctrl.AllRestaurants},
		{Method: fiber.MethodGet, Path: v1 + "/:id", Handler: ctrl.FindRestaurantById},
		{Method: fiber.MethodGet, Path: v1 + "/search", Handler: ctrl.SearchRestaurantByName},
		{
			Method:  fiber.MethodPost,
			Path:    v1 + "/create",
			Handler: ctrl.CreateRestaurant,
			Middlewares: []fiber.Handler{
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},
		{
			Method:  fiber.MethodPut,
			Path:    v1 + "/:id",
			Handler: ctrl.EditRestaurant,
			Middlewares: []fiber.Handler{
				middleware.Role(middleware.AllowedRoles{entity.RoleOwner}),
			},
		},
		{
			Method:  fiber.MethodDelete,
			Path:    v1 + "/:id",
			Handler: ctrl.DeleteRestaurant,
			Middlewares: []fiber.Handler{
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
