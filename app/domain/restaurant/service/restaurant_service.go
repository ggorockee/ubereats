package service

import (
	"ubereats/app/core/entity"

	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantService interface {
	GetAllRestaurant(context ...*fiber.Ctx) (*[]entity.Restaurant, error)
	CreateRestaurant(input *restaurantDto.CreateRestaurant, context ...*fiber.Ctx) (*entity.Restaurant, error)
	UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, context ...*fiber.Ctx) (*entity.Restaurant, error)
}

type restaurantService struct {
	db             *gorm.DB
	restaurantRepo restaurantRepo.RestaurantRepository
}

// UpdateRestaurant implements RestaurantService.
func (s *restaurantService) UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, context ...*fiber.Ctx) (*entity.Restaurant, error) {
	var err error
	var restaurant *entity.Restaurant
	err = s.db.Transaction(func(tx *gorm.DB) error {
		restaurant, err = s.restaurantRepo.UpdateRestaurant(input, id, context...)
		if err != nil {
			return err
		}
		return nil
	})
	return restaurant, err

}

// CreateRestaurant implements RestaurantService.
func (s *restaurantService) CreateRestaurant(input *restaurantDto.CreateRestaurant, context ...*fiber.Ctx) (*entity.Restaurant, error) {
	var err error
	var restaurant *entity.Restaurant
	err = s.db.Transaction(func(tx *gorm.DB) error {
		restaurant, err = s.restaurantRepo.CreateRestaurant(input, context...)
		if err != nil {
			return err
		}
		return nil
	})
	return restaurant, err
}

// GetAll implements RestaurantService.
func (s *restaurantService) GetAllRestaurant(context ...*fiber.Ctx) (*[]entity.Restaurant, error) {
	var err error
	var restaurants *[]entity.Restaurant

	err = s.db.Transaction(func(tx *gorm.DB) error {
		restaurants, err = s.restaurantRepo.GetAllRestaurant(context...)
		if err != nil {
			return err
		}
		return nil
	})

	return restaurants, err
}

func NewRestaurantService(d *gorm.DB, r restaurantRepo.RestaurantRepository) RestaurantService {
	return &restaurantService{
		db:             d,
		restaurantRepo: r,
	}
}
