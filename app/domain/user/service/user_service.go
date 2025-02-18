package service

import (
	"ubereats/app/core/entity"

	userDto "ubereats/app/domain/user/dto"
	userRepo "ubereats/app/domain/user/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUser(context ...*fiber.Ctx) (*[]entity.User, error)
	CreateAccount(input *userDto.CreateAccount, context ...*fiber.Ctx) (*entity.User, error)
	// UpdateUser(input *userDto.UpdateUser, id int, context ...*fiber.Ctx) (*entity.User, error)
}

type userService struct {
	db       *gorm.DB
	userRepo userRepo.UserRepository
}

// UpdateUser implements UserService.
// func (s *userService) UpdateUser(input *userDto.UpdateUser, id int, context ...*fiber.Ctx) (*entity.User, error) {
// 	var err error
// 	var user *entity.User
// 	err = s.db.Transaction(func(tx *gorm.DB) error {
// 		user, err = s.userRepo.UpdateUser(input, id, context...)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	return user, err

// }

// CreateUser implements UserService.
func (s *userService) CreateAccount(input *userDto.CreateAccount, context ...*fiber.Ctx) (*entity.User, error) {
	var err error
	var user *entity.User
	err = s.db.Transaction(func(tx *gorm.DB) error {
		user, err = s.userRepo.CreateAccount(input, context...)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

// GetAll implements UserService.
func (s *userService) GetAllUser(context ...*fiber.Ctx) (*[]entity.User, error) {
	var err error
	var users *[]entity.User

	err = s.db.Transaction(func(tx *gorm.DB) error {
		users, err = s.userRepo.GetAllUser(context...)
		if err != nil {
			return err
		}
		return nil
	})

	return users, err
}

func NewUserService(d *gorm.DB, r userRepo.UserRepository) UserService {
	return &userService{
		db:       d,
		userRepo: r,
	}
}
