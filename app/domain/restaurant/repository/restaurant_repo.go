package repository

import (
	"strings"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"

	"gorm.io/gorm"
)

type RestaurantRepository interface {
	FindByID(id string) (*entity.Restaurant, error)
	Create(param *restaurantDto.CreateRestaurantInput) (*entity.Restaurant, error)
	Update(param *restaurantDto.EditRestaurantInput) (*entity.Restaurant, error)
	Delete(id string) error
	FindAll(page int) (*[]entity.Restaurant, *int, error)
	SearchByName(query string, page int) (*[]entity.Restaurant, *int, error)
	CountByCategory(category *entity.Category) (*int, error)
}

func NewRestaurantRepository(d *gorm.DB) RestaurantRepository {
	return &restaurantRepository{
		dbConn: d,
	}
}

type restaurantRepository struct {
	dbConn *gorm.DB
}

// Create implements RestaurantRepository.
func (r *restaurantRepository) Create(param *restaurantDto.CreateRestaurantInput) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant
	if err := common.DecodeStructure(param, &restaurant); err != nil {
		return nil, err
	}

	if err := r.dbConn.Create(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

// CountByCategory implements RestaurantRepository.
func (r *restaurantRepository) CountByCategory(category *entity.Category) (*int, error) {
	var count int64
	err := r.dbConn.Model(&entity.Restaurant{}).
		Where("category_id = ?", category.ID).
		Count(&count).
		Error
	if err != nil {
		return nil, err
	}

	resultCount := int(count)
	return &resultCount, nil
}

// Delete implements RestaurantRepository.
func (r *restaurantRepository) Delete(id string) error {
	restaurant, err := r.FindByID(id)
	if err != nil {
		return err
	}

	err = r.dbConn.Delete(restaurant).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements RestaurantRepository.
func (r *restaurantRepository) FindAll(page int) (*[]entity.Restaurant, *int, error) {
	var restaurants []entity.Restaurant
	offset := (page - 1) * 25
	var total int64

	err := r.dbConn.Model(&entity.Restaurant{}).Count(&total).Error
	if err != nil {
		return nil, nil, err
	}

	err = r.dbConn.Limit(25).Offset(offset).Find(&restaurants).Error
	if err != nil {
		return nil, nil, err
	}

	totalCount := int(total)
	return &restaurants, &totalCount, nil
}

// FindByID implements RestaurantRepository.
func (r *restaurantRepository) FindByID(id string) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant
	err := r.dbConn.Where("id = ?", id).First(&restaurant).Error
	if err != nil {
		return nil, err
	}

	return &restaurant, nil
}

// SearchByName implements RestaurantRepository.
func (r *restaurantRepository) SearchByName(query string, page int) (*[]entity.Restaurant, *int, error) {
	var restaurants []entity.Restaurant
	offset := (page - 1) * 25
	var total int64

	searchQuery := "%" + strings.ToLower(query) + "%"
	err := r.dbConn.Model(&entity.Restaurant{}).
		Where("lower(name) LIKE ?", searchQuery).
		Count(&total).
		Error
	if err != nil {
		return nil, nil, err
	}

	err = r.dbConn.Where("lower(name) LIKE ?", searchQuery).
		Limit(25).
		Offset(offset).
		Find(&restaurants).
		Error

	if err != nil {
		return nil, nil, err
	}

	totalCount := int(total)

	return &restaurants, &totalCount, nil
}

// Update implements RestaurantRepository.
func (r *restaurantRepository) Update(param *restaurantDto.EditRestaurantInput) (*entity.Restaurant, error) {
	restaurant, err := r.FindByID(param.RestaurantID)
	if err != nil {
		return nil, err
	}

	if err := common.DecodeStructure(param, restaurant); err != nil {
		return nil, err
	}

	if err := r.dbConn.Model(&entity.Restaurant{}).Save(restaurant).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}
