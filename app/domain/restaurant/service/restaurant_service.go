package service

import (
	"fmt"
	"ubereats/app/core/entity"

	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantRes "ubereats/app/domain/restaurant/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantService interface {
	CreateRestaurant(input *restaurantDto.CreateRestaurant, c *fiber.Ctx) (*entity.Restaurant, error)
	UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, c *fiber.Ctx) (*entity.Restaurant, error)
	DeleteRestaurant(id int, c *fiber.Ctx) error
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
	var categories *[]entity.Category

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		categories, err = s.restaurantRepo.AllCategories()
		if err != nil {
			return err
		}
		return nil
	})

	return &restaurantRes.AllCategoriesResponse{
		BaseResponse: restaurantRes.BaseResponse{Ok: true},
		Categories:   *categories,
	}, err
}

// AllRestaurants implements RestaurantService.
func (s *restaurantService) AllRestaurants(input restaurantDto.RestaurantsInput) (*restaurantRes.RestaurantsResponse, error) {
	var restaurants *[]entity.Restaurant
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurants, err = s.restaurantRepo.AllRestaurants(input)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &restaurantRes.RestaurantsResponse{
			BaseResponse: restaurantRes.BaseResponse{
				Ok:    false,
				Error: "Could not load AllRestaurant",
			},
		}, err
	}

	var totalResults int64
	err = s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Restaurant{}).Count(&totalResults).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
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
		Results:      *restaurants,
		TotalPages:   int(totalResults+24) / 25, // ceil 처리
		TotalResults: int(totalResults),
	}, err
}

// CountRestaurants implements RestaurantService.
func (s *restaurantService) CountRestaurants(category entity.Category) (int, error) {
	var count int
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		count, err = s.restaurantRepo.CountRestaurants(category)
		if err != nil {
			return err
		}
		return nil
	})
	return count, err
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
	user, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return fmt.Errorf("user not authenticated")
	}

	restaurant, err := s.restaurantRepo.FindRestaurantById(restaurantDto.RestaurantInput{RestaurantID: id})
	if err != nil {
		return err
	}
	if restaurant.OwnerID != user.ID {
		return fmt.Errorf("you can't delete a restaurant you don't own")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.restaurantRepo.DeleteRestaurant(id, c)
	})
}

// FindRestaurantById implements RestaurantService.
func (s *restaurantService) FindRestaurantById(input restaurantDto.RestaurantInput) (*restaurantRes.RestaurantResponse, error) {
	var restaurant *entity.Restaurant
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.FindRestaurantById(input)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.RestaurantResponse{
			BaseResponse: restaurantRes.BaseResponse{
				Ok:    false,
				Error: "Restaurant not found",
			},
		}, err
	}

	return &restaurantRes.RestaurantResponse{
		BaseResponse: restaurantRes.BaseResponse{Ok: true},
		Restaurant:   restaurant,
	}, nil
}

// SearchRestaurantByName implements RestaurantService.
func (s *restaurantService) SearchRestaurantByName(input restaurantDto.SearchRestaurant) (*restaurantRes.SearchRestaurantResponse, error) {
	var restaurants *[]entity.Restaurant
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurants, err = s.restaurantRepo.SearchRestaurantByName(input)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.SearchRestaurantResponse{
			BaseResponse: restaurantRes.BaseResponse{Ok: true, Error: "Could not search for restaurants"},
		}, err
	}

	var totalResults int64
	err = s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Restaurant{}).Where("name LIKE ?", "%"+input.Query+"%").Count(&totalResults).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.SearchRestaurantResponse{
			BaseResponse: restaurantRes.BaseResponse{Ok: false, Error: "Could not count restaurants"},
			Restaurants:  *restaurants,
			TotalPages:   0,
			TotalResults: 0,
		}, err
	}

	return &restaurantRes.SearchRestaurantResponse{
		BaseResponse: restaurantRes.BaseResponse{Ok: true},
		Restaurants:  *restaurants,
		TotalPages:   int(totalResults+24) / 25,
		TotalResults: int(totalResults),
	}, nil
}

// UpdateRestaurant implements RestaurantService.
func (s *restaurantService) UpdateRestaurant(input *restaurantDto.UpdateRestaurant, id int, c *fiber.Ctx) (*entity.Restaurant, error) {
	user, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}

	restaurant, err := s.restaurantRepo.FindRestaurantById(restaurantDto.RestaurantInput{RestaurantID: id})
	if err != nil {
		return nil, err
	}

	if restaurant.OwnerID != user.ID {
		return nil, fmt.Errorf("you can't update a restaurant you don't own")
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
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
