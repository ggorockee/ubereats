package service

import (
	userRepo "ubereats/app/domain/user/repository"

	"gorm.io/gorm"
)

type UserService interface{}

type userService struct {
	dbConn   *gorm.DB
	userRepo userRepo.UserRepository
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
