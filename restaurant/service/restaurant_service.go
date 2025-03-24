package service

import (
	"github.com/gofiber/fiber/v2"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/response"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantRes "ubereats/app/domain/restaurant/response"

	"gorm.io/gorm"
)

type RestaurantService interface {
	CreateRestaurant(input *restaurantDto.CreateRestaurantInput, c *fiber.Ctx) (*restaurantRes.CreateRestaurantOutput, error)
	EditRestaurant(input *restaurantDto.EditRestaurantInput, c *fiber.Ctx) (*restaurantRes.EditRestaurantOutput, error)
	DeleteRestaurant(input *restaurantDto.DeleteRestaurantInput, c *fiber.Ctx) (*restaurantRes.DeleteRestaurantOutput, error)
	SearchRestaurantByName(input *restaurantDto.SearchRestaurantInput) (*restaurantRes.SearchRestaurantOutput, error)
	FindRestaurantById(input *restaurantDto.RestaurantInput) (*restaurantRes.RestaurantOutput, error)
	AllRestaurants(input *restaurantDto.RestaurantsInput) (*restaurantRes.RestaurantsOutput, error)
}

type restaurantService struct {
	dbConn         *gorm.DB
	restaurantRepo restaurantRepo.RestaurantRepository
}

func (s *restaurantService) CreateRestaurant(input *restaurantDto.CreateRestaurantInput, c *fiber.Ctx) (*restaurantRes.CreateRestaurantOutput, error) {
	var restaurant *entity.Restaurant
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.CreateRestaurant(input, c)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.CreateRestaurantOutput{
			BaseResponse: response.BaseResponse{Message: "Could not create restaurant"},
		}, err
	}

	return &restaurantRes.CreateRestaurantOutput{
		BaseResponse: response.BaseResponse{Ok: true},
		Restaurant: &restaurantRes.Restaurant{
			Name:     restaurant.Name,
			CoverImg: restaurant.CoverImg,
			Address:  restaurant.Address,
		},
	}, nil
}

func (s *restaurantService) EditRestaurant(input *restaurantDto.EditRestaurantInput, c *fiber.Ctx) (*restaurantRes.EditRestaurantOutput, error) {
	var restaurant *entity.Restaurant
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.EditRestaurant(input, c)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &restaurantRes.EditRestaurantOutput{
			BaseResponse: response.BaseResponse{Message: err.Error()},
		}, err
	}

	return &restaurantRes.EditRestaurantOutput{
		BaseResponse: response.BaseResponse{Ok: true},
		Restaurant: &restaurantRes.Restaurant{
			Name:     restaurant.Name,
			CoverImg: restaurant.CoverImg,
			Address:  restaurant.Address,
		},
	}, nil
}

func (s *restaurantService) DeleteRestaurant(input *restaurantDto.DeleteRestaurantInput, c *fiber.Ctx) (*restaurantRes.DeleteRestaurantOutput, error) {
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		err := s.restaurantRepo.DeleteRestaurant(input.RestaurantID, c)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.DeleteRestaurantOutput{
			BaseResponse: response.BaseResponse{Message: err.Error()},
		}, err
	}

	return &restaurantRes.DeleteRestaurantOutput{
		BaseResponse: response.BaseResponse{Ok: true},
	}, nil
}

func (s *restaurantService) SearchRestaurantByName(input *restaurantDto.SearchRestaurantInput) (*restaurantRes.SearchRestaurantOutput, error) {
	var restaurants *[]entity.Restaurant
	var total *int
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurants, total, err = s.restaurantRepo.FindByName(input.Query, input.Page)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.SearchRestaurantOutput{
			BaseResponse: response.BaseResponse{Message: err.Error()},
		}, err
	}

	return &restaurantRes.SearchRestaurantOutput{
		BaseResponse: response.BaseResponse{Ok: true},
		Restaurants:  nil,
		TotalResults: *total,
		TotalPages:   0,
	}, nil
}

func (s *restaurantService) FindRestaurantById(input *restaurantDto.RestaurantInput) (*restaurantRes.RestaurantOutput, error) {
	var restaurant *entity.Restaurant
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.FindRestaurantById(input.RestaurantID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &restaurantRes.RestaurantOutput{
			BaseResponse: response.BaseResponse{Message: "Could not create restaurant"},
		}, err
	}

	return &restaurantRes.RestaurantOutput{
		BaseResponse: response.BaseResponse{Ok: true},
	}, nil
}

func (s *restaurantService) AllRestaurants(input *restaurantDto.RestaurantsInput) (*restaurantRes.RestaurantsOutput, error) {
	var restaurants []entity.Restaurant
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var total *int
		var err error
		restaurants, total, err = s.restaurantRepo.FindAll(input.Page)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.RestaurantsOutput{
			BaseResponse: response.BaseResponse{Message: err.Error()},
		}, err
	}

	return &restaurantRes.RestaurantsOutput{
		BaseResponse: response.BaseResponse{Ok: true},
		Results:      nil,
	}, nil
}

// AllCategories implements RestaurantService.

func NewRestaurantService(d *gorm.DB, r restaurantRepo.RestaurantRepository) RestaurantService {
	return &restaurantService{
		dbConn:         d,
		restaurantRepo: r,
	}
}
