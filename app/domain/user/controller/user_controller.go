package controller

import (
	"fmt"
	"ubereats/app"
	"ubereats/app/core/helper/common"
	userDto "ubereats/app/domain/user/dto"
	userRes "ubereats/app/domain/user/response"
	userSvc "ubereats/app/domain/user/service"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Table() []app.Mapping
	GetAllUser(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
	// UpdateUser(c *fiber.Ctx) error
}

type userController struct {
	userSvc userSvc.UserService
}

// UpdateUser implements UserController.
// func (ctrl *userController) UpdateUser(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		common.ErrorResponse(c, common.ErrArg{
// 			IsError: true,
// 			Code:    fiber.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 	}

// 	var requestBody userDto.UpdateUser
// 	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
// 		return common.ErrorResponse(c, common.ErrArg{
// 			IsError: true,
// 			Code:    fiber.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 	}

// 	user, err := ctrl.userSvc.UpdateUser(&requestBody, id, c)
// 	if err != nil {
// 		return common.ErrorResponse(c, common.ErrArg{
// 			IsError: true,
// 			Code:    fiber.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 	}

// 	userResponse := userRes.GenUserRes(user)
// 	return common.SuccessResponse(c, common.SuccessArg{
// 		Message: "Success",
// 		Data:    userResponse,
// 	})
// }

func (ctrl *userController) CreateAccount(c *fiber.Ctx) error {
	var requestBody userDto.CreateAccount
	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if !requestBody.Role.IsValid() {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Invalid role: %s", requestBody.Role),
			Data:    nil,
		})
	}

	user, err := ctrl.userSvc.CreateAccount(&requestBody, c)
	if err != nil {
		return common.ErrorResponse(c, common.ErrArg{
			IsError: true,
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userResponse := userRes.GenUserRes(user)
	return common.SuccessResponse(c, common.SuccessArg{
		Message: "Success",
		Data:    userResponse,
	})

}

// GetAll implements UserController.
func (ctrl *userController) GetAllUser(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Table implements UserController.
func (ctrl *userController) Table() []app.Mapping {
	v1 := "/api/v1/user/create"

	return []app.Mapping{
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "",
			Handler: ctrl.GetAllUser,
		},

		{
			Method:  fiber.MethodPost,
			Path:    v1 + "",
			Handler: ctrl.CreateAccount,
		},

		// {
		// 	Method:  fiber.MethodPut,
		// 	Path:    v1 + "/:id",
		// 	Handler: ctrl.UpdateUser,
		// },
	}
}

func NewUserController(r userSvc.UserService) UserController {
	return &userController{
		userSvc: r,
	}
}
