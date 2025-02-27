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
	CreateRestaurant(input restaurantDto.CreateRestaurantInput, c *fiber.Ctx) (*restaurantRes.RestaurantOutput, error)
	EditRestaurant(ownerID string, input restaurantDto.EditRestaurantInput) (*restaurantRes.EditRestaurantOutput, error)
	DeleteRestaurant(ownerID string, input restaurantDto.DeleteRestaurantInput) (*restaurantRes.DeleteRestaurantOutput, error)
	AllRestaurants(input restaurantDto.RestaurantsInput) (*restaurantRes.RestaurantsOutput, error)
	FindRestaurantById(input restaurantDto.RestaurantInput) (*restaurantRes.RestaurantOutput, error)
	SearchRestaurantByName(input restaurantDto.SearchRestaurantInput) (*restaurantRes.SearchRestaurantOutput, error)
	CreateDish(ownerID string, input restaurantDto.CreateDishInput) (*restaurantRes.CreateDishOutput, error)
	EditDish(ownerID string, input restaurantDto.EditDishInput) (*restaurantRes.EditDishOutput, error)
	DeleteDish(ownerID string, input restaurantDto.DeleteDishInput) (*restaurantRes.DeleteDishOutput, error)
	AllCategories() (*restaurantRes.AllCategoriesOutput, error)
	FindCategoryBySlug(input restaurantDto.CategoryInput) (*restaurantRes.CategoryOutput, error)
}

type restaurantService struct {
	dbConn         *gorm.DB
	restaurantRepo restaurantRepo.RestaurantRepository
	categoryRepo   restaurantRepo.CategoryRepository
}

// AllCategories implements RestaurantService.
func (s *restaurantService) AllCategories() (*restaurantRes.AllCategoriesOutput, error) {
	var categories *[]entity.Category
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		categories, err = s.categoryRepo.FindAll()
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.AllCategoriesOutput{
			BaseResponse: restaurantRes.BaseResponse{
				Ok:    false,
				Error: "Could not load category restaurants",
			},
		}, err
	}

	return &restaurantRes.AllCategoriesOutput{
		BaseResponse: restaurantRes.BaseResponse{
			Ok: true,
		},
		Categories: restaurantRes.ResponseCategoriesOutput(categories),
	}, nil
}

// AllRestaurants implements RestaurantService.
func (s *restaurantService) AllRestaurants(input restaurantDto.RestaurantsInput) (*restaurantRes.RestaurantsOutput, error) {
	var restaurants *[]entity.Restaurant
	var total *int
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurants, total, err = s.restaurantRepo.FindAll(input.Page)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &restaurantRes.RestaurantsOutput{
			BaseResponse: restaurantRes.BaseResponse{Ok: false, Error: "Restaurant not found"},
		}, err
	}

	return &restaurantRes.RestaurantsOutput{
		BaseResponse: restaurantRes.BaseResponse{Ok: true},
		Results:      restaurants,
		TotalPages:   (*total + 24) / 25,
		TotalResults: *total,
	}, nil
}

// CreateDish implements RestaurantService.
func (s *restaurantService) CreateDish(ownerID string, input restaurantDto.CreateDishInput) (*restaurantRes.CreateDishOutput, error) {
	panic("unimplemented")
}

// CreateRestaurant implements RestaurantService.
func (s *restaurantService) CreateRestaurant(input *restaurantDto.CreateRestaurantInput, c *fiber.Ctx) (*entity.RestaurantResponse, error) {
	var restaurant *entity.Restaurant
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		restaurant, err = s.restaurantRepo.CreateRestaurant(input, c)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &entity.RestaurantResponse{
			CoreResponse: entity.CoreResponse{
				Ok:      false,
				Message: err.Error(),
			},
		}, err
	}

	return &entity.RestaurantResponse{
		CoreResponse: entity.CoreResponse{
			Ok: true,
		},
		Name:     restaurant.Name,
		CoverImg: restaurant.CoverImg,
		Address:  restaurant.Address,
		Category: restaurant.Serialize().Category,
		Owner:    restaurant.Serialize().Owner,
		Menu:     restaurant.Serialize().Menu,
	}, nil
}

// DeleteDish implements RestaurantService.
func (s *restaurantService) DeleteDish(ownerID string, input restaurantDto.DeleteDishInput) (*restaurantRes.DeleteDishOutput, error) {
	panic("unimplemented")
}

// DeleteRestaurant implements RestaurantService.
func (s *restaurantService) DeleteRestaurant(ownerID string, input restaurantDto.DeleteRestaurantInput) (*restaurantRes.DeleteRestaurantOutput, error) {
	panic("unimplemented")
}

// EditDish implements RestaurantService.
func (s *restaurantService) EditDish(ownerID string, input restaurantDto.EditDishInput) (*restaurantRes.EditDishOutput, error) {
	panic("unimplemented")
}

// EditRestaurant implements RestaurantService.
func (s *restaurantService) EditRestaurant(ownerID string, input restaurantDto.EditRestaurantInput) (*restaurantRes.EditRestaurantOutput, error) {
	panic("unimplemented")
}

// FindCategoryBySlug implements RestaurantService.
func (s *restaurantService) FindCategoryBySlug(input restaurantDto.CategoryInput) (*restaurantRes.CategoryOutput, error) {
	panic("unimplemented")
}

// FindRestaurantById implements RestaurantService.
func (s *restaurantService) FindRestaurantById(input restaurantDto.RestaurantInput) (*restaurantRes.RestaurantOutput, error) {
	panic("unimplemented")
}

// SearchRestaurantByName implements RestaurantService.
func (s *restaurantService) SearchRestaurantByName(input restaurantDto.SearchRestaurantInput) (*restaurantRes.SearchRestaurantOutput, error) {
	panic("unimplemented")
}

func NewRestaurantService(d *gorm.DB, r restaurantRepo.RestaurantRepository, c restaurantRepo.CategoryRepository) RestaurantService {
	return &restaurantService{
		dbConn:         d,
		restaurantRepo: r,
		categoryRepo:   c,
	}
}
