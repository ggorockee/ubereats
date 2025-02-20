package dish

import (
	"ubereats/app"
	dishCtrl "ubereats/app/domain/dish/controller"
	dishRepo "ubereats/app/domain/dish/repository"
	dishSvc "ubereats/app/domain/dish/service"

	"go.uber.org/fx"
)

var ControllerMoudle = fx.Module(
	"Controller",
	fx.Provide(
		dishRepo.NewDishRepository,
		dishSvc.NewDishService,
		app.AsRoute(dishCtrl.NewDishController),
	),
)
