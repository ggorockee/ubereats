package repository

import (
	"gorm.io/gorm"
)

type RestaurantRepository interface {
}

type restaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(d *gorm.DB) RestaurantRepository {
	return &restaurantRepository{
		db: d,
	}
}
