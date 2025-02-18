package user

import (
	"ubereats/app"
	userCtrl "ubereats/app/domain/user/controller"
	userRepo "ubereats/app/domain/user/repository"
	userSvc "ubereats/app/domain/user/service"

	"go.uber.org/fx"
)

var ControllerMoudle = fx.Module(
	"Controller",
	fx.Provide(
		userRepo.NewUserRepository,
		userSvc.NewUserService,
		app.AsRoute(userCtrl.NewUserController),
	),
)
