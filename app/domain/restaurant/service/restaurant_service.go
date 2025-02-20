package service

import (
	restaurantRepo "ubereats/app/domain/restaurant/repository"

	"gorm.io/gorm"
)

type RestaurantService interface {
}

type restaurantService struct {
	db             *gorm.DB
	restaurantRepo restaurantRepo.RestaurantRepository
}

func NewRestaurantService(d *gorm.DB, r restaurantRepo.RestaurantRepository) RestaurantService {
	return &restaurantService{
		db:             d,
		restaurantRepo: r,
	}
}
