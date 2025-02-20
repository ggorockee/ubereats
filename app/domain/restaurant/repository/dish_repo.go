package repository

import (
	"gorm.io/gorm"
)

type DishRepository interface {
}

type dishRepository struct {
	db *gorm.DB
}

func NewDishRepository(d *gorm.DB) DishRepository {
	return &dishRepository{
		db: d,
	}
}
