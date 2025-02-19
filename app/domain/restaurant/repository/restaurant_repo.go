package repository

import (
	"errors"
	"fmt"
	"log"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"

	restaurantDto "ubereats/app/domain/restaurant/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	GetAllRestaurant(context ...*fiber.Ctx) (*[]entity.Restaurant, error)
	CreateRestaurant(input *restaurantDto.CreateRestaurant, c *fiber.Ctx) (*entity.Restaurant, error)
	UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, c *fiber.Ctx) (*entity.Restaurant, error)
	GetFindById(id int, context ...*fiber.Ctx) (*entity.Restaurant, error)
}

type restaurantRepository struct {
	db *gorm.DB
}

// GetFindById implements RestaurantRepository.
func (r *restaurantRepository) GetFindById(id int, context ...*fiber.Ctx) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant
	if err := r.db.Preload("Owner").Preload("Category").Where("id = ?", id).First(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

// UpdateRestaurant implements RestaurantRepository.
func (r *restaurantRepository) UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, c *fiber.Ctx) (*entity.Restaurant, error) {
	restaurant, err := r.GetFindById(id, c)
	if err != nil {
		return nil, err
	}

	log.Println("categoryID >>>", input.CategoryID)
	if err := common.DecodeStructure(input, restaurant); err != nil {
		return nil, err
	}

	log.Println("categoryID >>>", restaurant.CategoryID)

	if err := restaurant.Validate(); err != nil {
		return nil, err
	}

	log.Println("categor~~~yID >>>", restaurant.CategoryID)

	// if err := r.db.Model(restaurant).Updates(restaurant).Error; err != nil {
	// 	return nil, err
	// }
	if err := r.db.Model(restaurant).Select("name", "cover_img", "address", "category_id", "owner_id").Updates(restaurant).Error; err != nil {
		return nil, fmt.Errorf("failed to update restaurant: %w", err)
	}

	log.Println("categoryID @@@", restaurant.CategoryID)

	if err := r.db.Preload("Owner").Preload("Category").Where("id = ?", id).First(&restaurant).Error; err != nil {
		return nil, err
	}

	log.Println("categoryID", restaurant.CategoryID)

	return restaurant, nil
}

// CreateRestaurant implements RestaurantRepository.
func (r *restaurantRepository) CreateRestaurant(input *restaurantDto.CreateRestaurant, c *fiber.Ctx) (*entity.Restaurant, error) {

	user, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, errors.New("user not authenticated")
	}

	var restaurant entity.Restaurant

	if err := common.DecodeStructure(input, &restaurant); err != nil {
		return nil, err
	}

	restaurant.OwnerID = user.ID

	if err := restaurant.Validate(); err != nil {
		return nil, err
	}

	if err := r.db.Create(&restaurant).Error; err != nil {
		return nil, err
	}

	if err := r.db.
		Preload("Category").
		Preload("Owner").
		First(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

// GetAll implements RestaurantRepository.
func (r *restaurantRepository) GetAllRestaurant(context ...*fiber.Ctx) (*[]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	if err := r.db.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return &restaurants, nil
}

func NewRestaurantRepository(d *gorm.DB) RestaurantRepository {
	return &restaurantRepository{db: d}
}
