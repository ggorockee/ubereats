package repository

import (
	"errors"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"

	categoryRepo "ubereats/app/domain/category/repository"
	restaurantDto "ubereats/app/domain/restaurant/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	CreateRestaurant(input *restaurantDto.CreateRestaurant, c *fiber.Ctx) (*entity.Restaurant, error)
	UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, c *fiber.Ctx) (*entity.Restaurant, error)
	DeleteRestaurant(id int, c *fiber.Ctx) error
	AllRestaurants(input restaurantDto.RestaurantsInput) (*[]entity.Restaurant, error)
	FindRestaurantById(input restaurantDto.RestaurantInput) (*entity.Restaurant, error)
	SearchRestaurantByName(input restaurantDto.SearchRestaurant) (*[]entity.Restaurant, error)
	AllCategories() (*[]entity.Category, error)
	CountRestaurants(category entity.Category) (int, error)
}

type restaurantRepository struct {
	db           *gorm.DB
	categoryRepo categoryRepo.CategoryRepository
}

func (r *restaurantRepository) SearchRestaurantByName(input restaurantDto.SearchRestaurant) (*[]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	if err := r.db.
		Where("name LIKE ?", "%"+input.Query+"%").
		Limit(25).
		Offset((input.Page - 1) * 25).
		Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return &restaurants, nil
}

func (r *restaurantRepository) FindRestaurantById(input restaurantDto.RestaurantInput) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant
	if err := r.db.First(&restaurant, input.RestaurantID).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (r *restaurantRepository) CountRestaurants(category entity.Category) (int, error) {
	var count int64
	if err := r.db.
		Model(&entity.Restaurant{}).
		Where("category_id = ?", category.ID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *restaurantRepository) AllRestaurants(input restaurantDto.RestaurantsInput) (*[]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	if err := r.db.
		Limit(25).
		Offset((input.Page - 1) * 25).
		Preload("Owner").
		Preload("Category").
		Find(&restaurants).Error; err != nil {
		return nil, errors.New("could not load restaurants")
	}

	return &restaurants, nil
}

func (r *restaurantRepository) AllCategories() (*[]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *restaurantRepository) DeleteRestaurant(id int, c *fiber.Ctx) error {
	restaurant, err := r.GetFindById(id, c)
	if err != nil {
		return err
	}

	if err := r.db.Delete(restaurant).Error; err != nil {
		return err
	}

	return nil
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

	if err := common.DecodeStructure(input, restaurant); err != nil {
		return nil, err
	}

	if err := restaurant.Validate(); err != nil {
		return nil, err
	}

	category, err := r.categoryRepo.GetFindById(restaurant.CategoryID)
	if err != nil {
		return nil, err
	}

	restaurant.Category = category

	if err := r.db.Save(restaurant).Error; err != nil {
		return nil, err
	}

	if err := r.db.Preload("Owner").Preload("Category").Where("id = ?", id).First(restaurant).Error; err != nil {
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

func NewRestaurantRepository(d *gorm.DB, c categoryRepo.CategoryRepository) RestaurantRepository {
	return &restaurantRepository{
		db:           d,
		categoryRepo: c,
	}
}
