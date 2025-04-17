package controller

import (
	"ubereats/app"
	userSvc "ubereats/app/domain/user/service"
	"ubereats/config"
)

type UserController interface {
	Table() []app.Mapping
}

type userController struct {
	userService userSvc.UserService
	cfg         *config.Config
}

func (ctrl *userController) Table() []app.Mapping {
	return []app.Mapping{}
}

func NewUserController(
	userService userSvc.UserService,
	cfg *config.Config,
) UserController {
	return &userController{
		userService: userService,
		cfg:         cfg,
	}
}
