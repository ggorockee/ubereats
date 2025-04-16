package repository

import (
	"ubereats/app/core/entity"
	restaurantDto "ubereats/app/domain/restaurant/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantRepo interface {
	CreateRestaurant(c *fiber.Ctx, inputParam *restaurantDto.CreateRestaurantDto) (*entity.Restaurant, error)
}

type restaurantRepo struct {
	dbConn *gorm.DB
}

func (r *restaurantRepo) CreateRestaurant(c *fiber.Ctx, inputParam *restaurantDto.CreateRestaurantDto) (*entity.Restaurant, error) {
	name := inputParam.Name
	addr := inputParam.Address
	// category := inputParam.Category

	restaurant := entity.Restaurant{
		Name:     name,
		Address:  addr,
		Category: &entity.Category{},
	}

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
