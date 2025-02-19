package category

import (
	"ubereats/app"
	categoryCtrl "ubereats/app/domain/category/controller"
	categoryRepo "ubereats/app/domain/category/repository"
	categorySvc "ubereats/app/domain/category/service"

	"go.uber.org/fx"
)

var ControllerMoudle = fx.Module(
	"Controller",
	fx.Provide(
		categoryRepo.NewCategoryRepository,
		categorySvc.NewCategoryService,
		app.AsRoute(categoryCtrl.NewCategoryController),
	),
)
