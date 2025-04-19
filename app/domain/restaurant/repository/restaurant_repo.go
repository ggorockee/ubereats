package repository

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	userRepo "ubereats/app/domain/user/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantRepo interface {
	CreateRestaurant(c *fiber.Ctx, inputParam *restaurantDto.CreateRestaurantIn) (*entity.Restaurant, error)
}

type restaurantRepo struct {
	dbConn   *gorm.DB
	userRepo userRepo.UserRepository
}

func (r *restaurantRepo) CreateRestaurant(c *fiber.Ctx, inputParam *restaurantDto.CreateRestaurantIn) (*entity.Restaurant, error) {
	// name := inputParam.Name
	// addr := inputParam.Address
	// category := inputParam.Category

	var restaurant entity.Restaurant
	authenticated_user, err := r.userRepo.GetAuthenticateUser(c)
	if err != nil {
		return nil, err
	}

	if err := common.DecodeStructure(inputParam, &restaurant); err != nil {
		return nil, err
	}

	// Owner 추가
	restaurant.Owner = *authenticated_user

	if err := r.dbConn.Create(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func NewRestaurantRepo(dbConn *gorm.DB) RestaurantRepo {
	return &restaurantRepo{
		dbConn: dbConn,
	}
}
