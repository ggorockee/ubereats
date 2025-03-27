package controller

// import (
// 	"ubereats/app"
// 	"ubereats/app/core/helper/common"
// 	userDto "ubereats/app/domain/user/dto"
// 	userSvc "ubereats/app/domain/user/service"
// 	"ubereats/app/middleware"
// 	"ubereats/config"

// 	"github.com/gofiber/fiber/v2"
// )

// type UserController interface {
// 	Table() []app.Mapping
// 	GetAllUser(c *fiber.Ctx) error
// 	CreateAccount(c *fiber.Ctx) error
// 	Login(c *fiber.Ctx) error
// 	Me(c *fiber.Ctx) error
// 	// UpdateUser(c *fiber.Ctx) error
// }

// type userController struct {
// 	userSvc userSvc.UserService
// 	cfg     *config.Config
// }

// // UpdateUser implements UserController.
// // func (ctrl *userController) UpdateUser(c *fiber.Ctx) error {
// // 	id, err := c.ParamsInt("id")
// // 	if err != nil {
// // 		common.ErrorResponse(c, common.ErrArg{
// // 			IsError: true,
// // 			Code:    fiber.StatusBadRequest,
// // 			Message: err.Error(),
// // 			Data:    nil,
// // 		})
// // 	}

// // 	var requestBody userDto.UpdateUser
// // 	if err := common.RequestParserAndValidate(c, &requestBody); err != nil {
// // 		return common.ErrorResponse(c, common.ErrArg{
// // 			IsError: true,
// // 			Code:    fiber.StatusBadRequest,
// // 			Message: err.Error(),
// // 			Data:    nil,
// // 		})
// // 	}

// // 	user, err := ctrl.userSvc.UpdateUser(&requestBody, id, c)
// // 	if err != nil {
// // 		return common.ErrorResponse(c, common.ErrArg{
// // 			IsError: true,
// // 			Code:    fiber.StatusBadRequest,
// // 			Message: err.Error(),
// // 			Data:    nil,
// // 		})
// // 	}

// // 	userResponse := userRes.GenUserRes(user)
// // 	return common.SuccessResponse(c, common.SuccessArg{
// // 		Message: "Success",
// // 		Data:    userResponse,
// // 	})
// // }

// // func (ctrl *userController) Me(c *fiber.Ctx) error {
// // 	user, err := ctrl.userSvc.Me(c)
// // 	if err != nil {
// // 		return common.BaseResponse{
// // 			Message: err.Error(),
// // 		}
// // 	}

// // 	return common.BaseResponse{
// // 		Message: err.Error(),
// // 	}

// // }

// func (ctrl *userController) Login(c *fiber.Ctx) error {
// 	var requestBody userDto.Login
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(
// 			common.CoreResponse{
// 				Message: err.Error(),
// 			},
// 		)
// 	}

// 	if err := common.ValidateStruct(&requestBody, c); err != nil {
// 		return err
// 	}

// 	output, err := ctrl.userSvc.Login(c, &requestBody)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(output)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(output)
// }

// func (ctrl *userController) CreateAccount(c *fiber.Ctx) error {
// 	var requestBody userDto.CreateAccount

// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(
// 			common.CoreResponse{
// 				Message: err.Error(),
// 			},
// 		)
// 	}

// 	if err := common.ValidateStruct(&requestBody, c); err != nil {
// 		return err
// 	}

// 	output, err := ctrl.userSvc.CreateAccount(c, &requestBody)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(output)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(output)
// }

// // GetAll implements UserController.
// func (ctrl *userController) GetAllUser(c *fiber.Ctx) error {
// 	panic("unimplemented")
// }

// // Table implements UserController.
// func (ctrl *userController) Table() []app.Mapping {
// 	v1 := "/api/v1/user"

// 	return []app.Mapping{
// 		{
// 			Method:  fiber.MethodGet,
// 			Path:    v1 + "",
// 			Handler: ctrl.GetAllUser,
// 		},

// 		{
// 			Method:  fiber.MethodPost,
// 			Path:    v1 + "/create",
// 			Handler: ctrl.CreateAccount,
// 		},

// 		{
// 			Method:  fiber.MethodPost,
// 			Path:    v1 + "/login",
// 			Handler: ctrl.Login,
// 		},

// 		{
// 			Method:  fiber.MethodGet,
// 			Path:    v1 + "/me",
// 			Handler: ctrl.Me,
// 			Middlewares: []fiber.Handler{
// 				middleware.AuthMiddleware(ctrl.cfg),
// 			},
// 		},

// 		// {
// 		// 	Method:  fiber.MethodPut,
// 		// 	Path:    v1 + "/:id",
// 		// 	Handler: ctrl.UpdateUser,
// 		// },
// 	}
// }

// func NewUserController(r userSvc.UserService, c *config.Config) UserController {
// 	return &userController{
// 		userSvc: r,
// 		cfg:     c,
// 	}
// }
