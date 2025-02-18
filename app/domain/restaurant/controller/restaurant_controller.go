package controller

import (
	restaurantSvc "ubereats/app/domain/restaurant/service"
)

type RestaurantController interface{}

type restaurantController struct {
	restaurantSvc restaurantSvc.RestaurantService
}

func NewRestaurantController(r restaurantSvc.RestaurantService) RestaurantController {
	return &restaurantController{
		restaurantSvc: r,
	}
}
