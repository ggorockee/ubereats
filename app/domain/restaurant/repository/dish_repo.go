package repository

import (
	restaurantDto "ubereats/app/domain/restaurant/dto"

	"gorm.io/gorm"
)

type DishRepository interface {
	CreateDish(dish *restaurantDto.Dish) error
	FindDishByID(id string) (*restaurantDto.Dish, error)
	UpdateDish(id string, dish *restaurantDto.Dish) error
	DeleteDish(id string) error
}

type dishRepository struct {
	db *gorm.DB
}

func NewDishRepository(d *gorm.DB) DishRepository {
	return &dishRepository{
		db: d,
	}
}
