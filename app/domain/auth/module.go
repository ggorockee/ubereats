package auth

import (
	"ubereats/app"
	authController "ubereats/app/domain/auth/controller"
	authRepo "ubereats/app/domain/auth/repository"
	authService "ubereats/app/domain/auth/service"

	"go.uber.org/fx"
)

var ControllerMoudle = fx.Module(
	"controller",
	fx.Provide(
		authService.NewAuthService,
		authRepo.NewAuthRepository,
		app.AsRoute(authController.NewAuthController),
	),
)
