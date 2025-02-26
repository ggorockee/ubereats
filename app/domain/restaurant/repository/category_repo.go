package repository

import (
	"errors"
	"fmt"
	"ubereats/app/core/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategoryByName(name string) (*entity.Category, error)
	FindAll() (*[]entity.Category, error)
}

type categoryRepository struct {
	dbConn *gorm.DB
}

// FindAll implements CategoryRepository.
func (r *categoryRepository) FindAll() (*[]entity.Category, error) {
	var categories []entity.Category
	if err := r.dbConn.Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

// FindCategoryByName implements CategoryRepository.
func (r *categoryRepository) FindCategoryByName(name string) (*entity.Category, error) {
	var category entity.Category
	if err := r.dbConn.Where("name = ?", name).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("cannot not found categoryf")
		}
		return nil, fmt.Errorf("DB Error")
	}

	return &category, nil
}

func NewCategoryRepository(d *gorm.DB) CategoryRepository {
	return &categoryRepository{
		dbConn: d,
	}
}
