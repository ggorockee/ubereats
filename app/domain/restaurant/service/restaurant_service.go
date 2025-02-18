package service

import (
	"ubereats/app/core/entity"

	restaurantRepo "ubereats/app/domain/restaurant/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantService interface {
	GetAll(context ...*fiber.Ctx) (*[]entity.Restaurant, error)
}

type restaurantService struct {
	db             *gorm.DB
	restaurantRepo restaurantRepo.RestaurantRepository
}

// GetAll implements RestaurantService.
func (s *restaurantService) GetAll(context ...*fiber.Ctx) (*[]entity.Restaurant, error) {
	var err error
	var restaurants *[]entity.Restaurant

	err = s.db.Transaction(func(tx *gorm.DB) error {
		restaurants, err = s.restaurantRepo.GetAll(context...)
		if err != nil {
			return err
		}
		return nil
	})

	return restaurants, err
}

func NewRestaurantService(d *gorm.DB) RestaurantService {
	return &restaurantService{
		db: d,
	}
}
