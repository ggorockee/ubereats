package restaurant

import (
	"ubereats/app"
	restaurantCtrl "ubereats/app/domain/restaurant/controller"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantSvc "ubereats/app/domain/restaurant/service"

	"go.uber.org/fx"
)

var ControllerMoudle = fx.Module(
	"Controller",
	fx.Provide(
		restaurantRepo.NewRestaurantRepository,
		restaurantSvc.NewRestaurantService,
		app.AsRoute(restaurantCtrl.NewRestaurantController),
	),
)
