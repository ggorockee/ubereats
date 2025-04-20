package repository

import (
	"fmt"
	"strconv"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	userRepo "ubereats/app/domain/user/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantRepo interface {
	CreateRestaurant(c *fiber.Ctx, inputParam *restaurantDto.CreateRestaurantIn) (*entity.Restaurant, error)
	GetAllRestaurant(c *fiber.Ctx) (*[]entity.Restaurant, error)
	EditRestaurant(c *fiber.Ctx, id uint, inputParam *restaurantDto.EditRestaurantIn) (*entity.Restaurant, error)
	FineOne(key, value string) (*entity.Restaurant, error)
}

type restaurantRepo struct {
	dbConn   *gorm.DB
	userRepo userRepo.UserRepository
	catRepo  CategoryRepository
}

// FineOne implements RestaurantRepo.
func (r *restaurantRepo) FineOne(key string, value string) (*entity.Restaurant, error) {
	var obj entity.Restaurant
	query := fmt.Sprintf("%s = ?", key)
	err := r.dbConn.
		Preload("Category").
		Preload("Owner").
		Where(query, value).First(&obj).Error
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

// EditRestaurant implements RestaurantRepo.
func (r *restaurantRepo) EditRestaurant(c *fiber.Ctx, id uint, inputParam *restaurantDto.EditRestaurantIn) (*entity.Restaurant, error) {
	restaurant, err := r.FineOne("id", fmt.Sprintf("%d", id))
	if err != nil {
		return nil, err
	}

	if err := common.DecodeStructure(inputParam, restaurant); err != nil {
		return nil, err
	}

	if err := common.ValidateStruct(restaurant); err != nil {
		return nil, err
	}

	// category setting
	if inputParam.CategoryId != nil {
		category, err := r.catRepo.FineOne("id", fmt.Sprintf("%d", *inputParam.CategoryId))
		if err != nil {
			return nil, err
		}

		restaurant.Category = *category
	}

	if err := r.dbConn.Save(restaurant).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}

// GetAllRestaurant implements RestaurantRepo.
func (r *restaurantRepo) GetAllRestaurant(c *fiber.Ctx) (*[]entity.Restaurant, error) {
	var restaurants []entity.Restaurant

	if err := r.dbConn.Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return &restaurants, nil
}

func (r *restaurantRepo) CreateRestaurant(c *fiber.Ctx, inputParam *restaurantDto.CreateRestaurantIn) (*entity.Restaurant, error) {
	// name := inputParam.Name
	// addr := inputParam.Address
	// category := inputParam.Category

	var restaurant entity.Restaurant
	authenticated_user, err := r.userRepo.GetAuthenticateUser(c)

	if err != nil {
		return nil, err
	}

	if err := common.DecodeStructure(inputParam, &restaurant); err != nil {
		return nil, err
	}
	// Owner 추가
	restaurant.Owner = *authenticated_user

	categoryIdStr := strconv.Itoa(int(inputParam.CategoryId))
	category, err := r.catRepo.FineOne("id", categoryIdStr)
	if err != nil {
		return nil, err
	}

	// Category
	restaurant.Category = *category
	if err := r.dbConn.Create(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func NewRestaurantRepo(
	dbConn *gorm.DB,
	userRepo userRepo.UserRepository,
	catRepo CategoryRepository,
) RestaurantRepo {
	return &restaurantRepo{
		dbConn:   dbConn,
		userRepo: userRepo,
		catRepo:  catRepo,
	}
}
