package service

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	authDto "ubereats/app/domain/auth/dto"
	authRepository "ubereats/app/domain/auth/repository"
	authResponse "ubereats/app/domain/auth/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthService interface {
	SignUp(inputParam *authDto.SignUpInput) (*authResponse.SignUpOutput, error)
	Login(c *fiber.Ctx, inputParam *authDto.LoginInput) (*authResponse.LoginOutput, error)
}

type authService struct {
	dbConn   *gorm.DB
	authRepo authRepository.AuthRepository
}

// Login implements AuthService.
func (s *authService) Login(c *fiber.Ctx, inputParam *authDto.LoginInput) (*authResponse.LoginOutput, error) {
	var loginUser *entity.User
	var token string

	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		token, err = s.authRepo.Login(c, inputParam)
		if err != nil {
			return err
		}

		loginUser, err = s.authRepo.FindByOne("email", inputParam.Email)
		if err != nil {
			return err
		}
		return nil
	})

	userResponse := struct {
		LoginUser entity.User `json:"login_user"`
		Token     string      `json:"token"`
	}{
		LoginUser: *loginUser,
		Token:     token,
	}

	if err != nil {
		return &authResponse.LoginOutput{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &authResponse.LoginOutput{
		CoreResponse: common.CoreResponse{
			Ok:      true,
			Message: "success login",
			Data:    userResponse,
		},
	}, nil
}

// SignUp implements AuthService.
func (s *authService) SignUp(inputParam *authDto.SignUpInput) (*authResponse.SignUpOutput, error) {
	var user *entity.User
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		user, err = s.authRepo.SignUp(inputParam)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &authResponse.SignUpOutput{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &authResponse.SignUpOutput{
		CoreResponse: common.CoreResponse{
			Ok:      true,
			Message: "success signup",
			Data:    user,
		},
	}, nil
}

func NewAuthService(
	dbConn *gorm.DB,
	authRepo authRepository.AuthRepository,
) AuthService {
	return &authService{
		dbConn:   dbConn,
		authRepo: authRepo,
	}
}
