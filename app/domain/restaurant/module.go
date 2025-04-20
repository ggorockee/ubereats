package restaurant

import (
	"ubereats/app"
	restaurantCtrl "ubereats/app/domain/restaurant/controller"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantSvc "ubereats/app/domain/restaurant/service"

	"go.uber.org/fx"
)

var ControllerModule = fx.Module(
	"Controller",
	fx.Provide(
		// category
		restaurantRepo.NewCategoryRepository,
		restaurantSvc.NewCategoryService,
		app.AsRoute(restaurantCtrl.NewCategoryContoller),

		// restaurant
		restaurantRepo.NewRestaurantRepo,
		restaurantSvc.NewRestaurantService,
		app.AsRoute(restaurantCtrl.NewRestaurantController),
	),
)
