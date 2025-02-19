package repository

import (
	"errors"
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
	if err := r.db.Where("id = ?", id).First(&restaurant).Error; err != nil {
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

	if err := common.DecodeStructure(input, restaurant); err != nil {
		return nil, err
	}

	if err := r.db.Save(restaurant).Error; err != nil {
		return nil, err
	}
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
