package repository

import (
	"ubereats/app/core/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	GetAll(context ...*fiber.Ctx) (*[]entity.Restaurant, error)
}

type restaurantRepository struct {
	db *gorm.DB
}

// GetAll implements RestaurantRepository.
func (r *restaurantRepository) GetAll(context ...*fiber.Ctx) (*[]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	if err := r.db.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return &restaurants, nil
}

func NewRestaurantRepository(d *gorm.DB) RestaurantRepository {
	return &restaurantRepository{db: d}
}
