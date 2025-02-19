package service

import (
	"ubereats/app/core/entity"

	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"

	restaurantRes "ubereats/app/domain/restaurant/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantService interface {
	GetAllRestaurant(c *fiber.Ctx) (*[]entity.Restaurant, error)
	CreateRestaurant(input *restaurantDto.CreateRestaurant, c *fiber.Ctx) (*entity.Restaurant, error)
	UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, c *fiber.Ctx) (*entity.Restaurant, error)
	DeleteRestaurant(id int, c *fiber.Ctx) error
	GetFindById(id int, c *fiber.Ctx) (*entity.Restaurant, error)
	AllRestaurants(input restaurantDto.RestaurantsInput) (*restaurantRes.RestaurantsResponse, error)
	FindRestaurantById(input restaurantDto.RestaurantInput) (*restaurantRes.RestaurantResponse, error)
	SearchRestaurantByName(input restaurantDto.SearchRestaurant) (*restaurantRes.SearchRestaurantResponse, error)
	AllCategories() (*restaurantRes.AllCategoriesResponse, error)
	CountRestaurants(category entity.Category) (int, error)
}

type restaurantService struct {
	db             *gorm.DB
	restaurantRepo restaurantRepo.RestaurantRepository
}

// AllCategories implements RestaurantService.
func (s *restaurantService) AllCategories() (*restaurantRes.AllCategoriesResponse, error) {
	var categories []entity.Category
	if err := s.db.Find(&categories).Error; err != nil {
		return &restaurantRes.AllCategoriesResponse{
			BaseResponse: restaurantRes.BaseResponse{Ok: false, Error: "Could not load categories"},
		}, err
	}

	return &restaurantRes.AllCategoriesResponse{
		BaseResponse: restaurantRes.BaseResponse{Ok: true},
		Categories:   categories,
	}, nil
}

// AllRestaurants implements RestaurantService.
func (s *restaurantService) AllRestaurants(input restaurantDto.RestaurantsInput) (*restaurantRes.RestaurantsResponse, error) {
	var restaurants []entity.Restaurant
	if err := s.db.Limit(25).
		Offset((input.Page - 1) * 25).
		Find(&restaurants).Error; err != nil {
		return &restaurantRes.RestaurantsResponse{
			BaseResponse: restaurantRes.BaseResponse{
				Ok:    false,
				Error: "Could not load restaurants",
			},
		}, err
	}

	var totalResults int64
	if err := s.db.Model(&entity.Restaurant{}).Count(&totalResults).Error; err != nil {
		return &restaurantRes.RestaurantsResponse{
			BaseResponse: restaurantRes.BaseResponse{
				Ok:    false,
				Error: "Could not load restaurants",
			},
		}, err
	}

	return &restaurantRes.RestaurantsResponse{
		BaseResponse: restaurantRes.BaseResponse{
			Ok: true,
		},
		Results:      restaurants,
		TotalPages:   int(totalResults+24) / 25, // ceil 처리
		TotalResults: int(totalResults),
	}, nil
}

// CountRestaurants implements RestaurantService.
func (s *restaurantService) CountRestaurants(category entity.Category) (int, error) {
	var count int64
	if err := s.db.Model(&entity.Restaurant{}).Where("category_id = ?", category.ID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// CreateRestaurant implements RestaurantService.
func (s *restaurantService) CreateRestaurant(input *restaurantDto.CreateRestaurant, c *fiber.Ctx) (*entity.Restaurant, error) {
	var restaurant *entity.Restaurant
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.CreateRestaurant(input, c)
		return err
	})
	return restaurant, err
}

// DeleteRestaurant implements RestaurantService.
func (s *restaurantService) DeleteRestaurant(id int, c *fiber.Ctx) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.restaurantRepo.DeleteRestaurant(id, c)
	})
}

// FindRestaurantById implements RestaurantService.
func (s *restaurantService) FindRestaurantById(input restaurantDto.RestaurantInput) (*restaurantRes.RestaurantResponse, error) {
	var restaurant entity.Restaurant
	if err := s.db.First(&restaurant, input.RestaurantID).Error; err != nil {
		return &restaurantRes.RestaurantResponse{
			BaseResponse: restaurantRes.BaseResponse{
				Ok:    false,
				Error: "Restaurant not found",
			},
		}, nil
	}
	return &restaurantRes.RestaurantResponse{
		BaseResponse: restaurantRes.BaseResponse{Ok: true},
		Restaurant:   &restaurant,
	}, nil

}

// GetAllRestaurant implements RestaurantService.
func (s *restaurantService) GetAllRestaurant(c *fiber.Ctx) (*[]entity.Restaurant, error) {
	var restaurants *[]entity.Restaurant
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurants, err = s.restaurantRepo.GetAllRestaurant(c)
		return err
	})
	return restaurants, err
}

// GetFindById implements RestaurantService.
func (s *restaurantService) GetFindById(id int, c *fiber.Ctx) (*entity.Restaurant, error) {
	var restaurant *entity.Restaurant
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.GetFindById(id, c)
		return err
	})
	return restaurant, err
}

// SearchRestaurantByName implements RestaurantService.
func (s *restaurantService) SearchRestaurantByName(input restaurantDto.SearchRestaurant) (*restaurantRes.SearchRestaurantResponse, error) {
	var restaurants []entity.Restaurant
	if err := s.db.Where("name LIKE ?", "%"+input.Query+"%").
		Limit(25).
		Offset((input.Page - 1) * 25).
		Find(&restaurants).Error; err != nil {
		return &restaurantRes.SearchRestaurantResponse{
			BaseResponse: restaurantRes.BaseResponse{Ok: true, Error: "Could not search for restaurants"},
		}, nil
	}

	var totalResults int64
	if err := s.db.Model(&entity.Restaurant{}).Where("name LIKE ?", "%"+input.Query+"%").Count(&totalResults).Error; err != nil {
		return &restaurantRes.SearchRestaurantResponse{
			BaseResponse: restaurantRes.BaseResponse{Ok: false, Error: "Could not count restaurants"},
			Restaurants:  restaurants,
			TotalPages:   0,
			TotalResults: 0,
		}, err
	}

	return &restaurantRes.SearchRestaurantResponse{
		BaseResponse: restaurantRes.BaseResponse{Ok: true},
		Restaurants:  restaurants,
		TotalPages:   int(totalResults+24) / 25,
		TotalResults: int(totalResults),
	}, nil
}

// UpdateRestaurant implements RestaurantService.
func (s *restaurantService) UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, c *fiber.Ctx) (*entity.Restaurant, error) {
	var restaurant *entity.Restaurant
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.UpdateRestaurant(input, id, c)
		return err
	})
	return restaurant, err
}

func NewRestaurantService(d *gorm.DB, r restaurantRepo.RestaurantRepository) RestaurantService {
	return &restaurantService{
		db:             d,
		restaurantRepo: r,
	}
}
