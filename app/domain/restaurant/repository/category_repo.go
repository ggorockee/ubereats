package repository

import (
	"fmt"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(inputParam *restaurantDto.CreateCategoryInput) (*entity.Category, error)
}

type categoryRepository struct {
	dbConn *gorm.DB
}

// CreateCategory implements CategoryRepository.
func (r *categoryRepository) CreateCategory(inputParam *restaurantDto.CreateCategoryInput) (*entity.Category, error) {
	var category entity.Category
	if err := common.DecodeStructure(inputParam, &category); err != nil {
		return nil, fmt.Errorf("decode Structure %w", err)
	}

	if err := r.dbConn.Create(&category).Error; err != nil {
		return nil, fmt.Errorf("create error %w", err)
	}

	return &category, nil
}

func NewCategoryRepository(dbConn *gorm.DB) CategoryRepository {
	return &categoryRepository{
		dbConn: dbConn,
	}
}
