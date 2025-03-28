package restaurant

import (
	"ubereats/app"
	restaurantController "ubereats/app/domain/restaurant/controller"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantService "ubereats/app/domain/restaurant/service"

	"go.uber.org/fx"
)

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		restaurantRepo.NewCategoryRepository,
		restaurantService.NewCategoryService,
		app.AsRoute(restaurantController.NewCategoryController),
	),
)
