package repository

import (
	"gorm.io/gorm"
)

type CategoryRepository interface {
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(d *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: d,
	}
}
