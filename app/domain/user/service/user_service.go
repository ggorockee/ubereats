package service

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	userDto "ubereats/app/domain/user/dto"
	userRepo "ubereats/app/domain/user/repository"
	userResp "ubereats/app/domain/user/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	CreateAccount(c *fiber.Ctx, inputParam *userDto.CreateAccountIn) (*userResp.CreateAccountOut, error)
}

type userService struct {
	dbConn   *gorm.DB
	userRepo userRepo.UserRepository
}

// CreateAccount implements UserService.
func (s *userService) CreateAccount(c *fiber.Ctx, inputParam *userDto.CreateAccountIn) (*userResp.CreateAccountOut, error) {
	var user *entity.User
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		user, err = s.userRepo.CreateAccount(c, inputParam)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &userResp.CreateAccountOut{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &userResp.CreateAccountOut{
		CoreResponse: common.CoreResponse{
			Ok:   true,
			Data: user,
		},
	}, nil
}

func NewUserService(
	dbConn *gorm.DB,
	userRepo userRepo.UserRepository,
) UserService {
	return &userService{
		dbConn:   dbConn,
		userRepo: userRepo,
	}
}
