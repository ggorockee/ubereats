package service

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantResp "ubereats/app/domain/restaurant/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantService interface {
	CreateRestaurant(c *fiber.Ctx, inputParm *restaurantDto.CreateRestaurantDto) (*restaurantResp.CreateRestaurantOut, error)
}

type restaurantService struct {
	restaurantRepo restaurantRepo.RestaurantRepo
	dbConn         *gorm.DB
}

// CreateRestaurant implements RestaurantService.
func (s *restaurantService) CreateRestaurant(c *fiber.Ctx, inputParm *restaurantDto.CreateRestaurantDto) (*restaurantResp.CreateRestaurantOut, error) {
	var restaurant *entity.Restaurant
	// var err error
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.CreateRestaurant(c, inputParm)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantResp.CreateRestaurantOut{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &restaurantResp.CreateRestaurantOut{
		CoreResponse: common.CoreResponse{
			Ok:   true,
			Data: restaurant,
		},
	}, nil
}

func NewRestaurantService(
	restaurantRepo restaurantRepo.RestaurantRepo,
	dbConn *gorm.DB,
) RestaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
		dbConn:         dbConn,
	}
}
