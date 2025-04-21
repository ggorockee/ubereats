package service

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantResp "ubereats/app/domain/restaurant/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestaurantService interface {
	CreateRestaurant(c *fiber.Ctx, inputParm *restaurantDto.CreateRestaurantIn) (*restaurantResp.CreateRestaurantOut, error)
	AllRestaurant(c *fiber.Ctx, param ...restaurantDto.GetRestaurantsInput) (*restaurantResp.AllRestaurantOut, error)
	EditRestaurant(c *fiber.Ctx, id uint, inputParam *restaurantDto.EditRestaurantIn) (*restaurantResp.EditRestaurantOut, error)
	FindCategoryByName(c *fiber.Ctx, name string, params ...common.PaginationParams) (*restaurantResp.GetCategoryOut, error)
}

type restaurantService struct {
	restaurantRepo restaurantRepo.RestaurantRepo
	dbConn         *gorm.DB
}

// FindCategoryByName implements RestaurantService.
func (s *restaurantService) FindCategoryByName(c *fiber.Ctx, name string, params ...common.PaginationParams) (*restaurantResp.GetCategoryOut, error) {
	p := common.PaginationParams{
		Page:  1,
		Limit: 25,
	}
	if len(params) > 0 {
		p = params[0] // 첫 번째 요소만 사용

		if p.Limit == 0 {
			p.Limit = 25
		}
	}

	category, err := s.restaurantRepo.FindCategoryByName(c, name, p)
	if err != nil {
		return &restaurantResp.GetCategoryOut{
			Message: err.Error(),
		}, nil
	}

	restaurants := restaurantResp.ToSimpleRestaurants(category.Restaurants)

	result := restaurantResp.CategoryResult{
		ID:          category.ID,
		Name:        category.Name,
		CoverImg:    category.CoverImg,
		Restaurants: restaurants,
		TotalPages:  *category.TotalPages,
	}

	return &restaurantResp.GetCategoryOut{
		Ok:   true,
		Data: &result,
		PaginationOutput: common.PaginationOutput{
			TotalPages: category.TotalPages,
		},
	}, nil
}

// EditRestaurant implements RestaurantService.
func (s *restaurantService) EditRestaurant(c *fiber.Ctx, id uint, inputParam *restaurantDto.EditRestaurantIn) (*restaurantResp.EditRestaurantOut, error) {
	var restaurant *entity.Restaurant
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.EditRestaurant(c, id, inputParam)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantResp.EditRestaurantOut{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &restaurantResp.EditRestaurantOut{
		CoreResponse: common.CoreResponse{
			Ok:   true,
			Data: restaurant,
		},
	}, nil
}

// GetAllRestaurant implements RestaurantService.
func (s *restaurantService) AllRestaurant(c *fiber.Ctx, param ...restaurantDto.GetRestaurantsInput) (*restaurantResp.AllRestaurantOut, error) {

	// 1. 기본값 설정
	p := restaurantDto.GetRestaurantsInput{
		Page:  1,
		Limit: 25,
	}

	if len(param) > 0 {
		p = param[0]
		if p.Limit == 0 {
			p.Limit = 25
		}
		if p.Page == 0 {
			p.Page = 1
		}
	}

	// 2. 레포지토리 호출
	restaurants, totalResults, err := s.restaurantRepo.AllRestaurant(c, p)
	if err != nil {
		return &restaurantResp.AllRestaurantOut{
			Message: "could not load restaurants",
		}, nil
	}

	// 3. 총 페이지 계산
	totalPages := (totalResults + (p.Limit) - 1) / (p.Limit)

	simpleRestaurants := restaurantResp.ToSimpleRestaurants(*restaurants)

	return &restaurantResp.AllRestaurantOut{
		Ok:           false,
		Message:      "",
		Results:      simpleRestaurants,
		TotalPages:   totalPages,
		TotalResults: totalResults,
	}, nil
}

// CreateRestaurant implements RestaurantService.
func (s *restaurantService) CreateRestaurant(c *fiber.Ctx, inputParm *restaurantDto.CreateRestaurantIn) (*restaurantResp.CreateRestaurantOut, error) {
	var restaurant *entity.Restaurant
	// var err error
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		restaurant, err = s.restaurantRepo.CreateRestaurant(c, inputParm)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantResp.CreateRestaurantOut{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &restaurantResp.CreateRestaurantOut{
		CoreResponse: common.CoreResponse{
			Ok:   true,
			Data: restaurant,
		},
	}, nil
}

func NewRestaurantService(
	restaurantRepo restaurantRepo.RestaurantRepo,
	dbConn *gorm.DB,
) RestaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
		dbConn:         dbConn,
	}
}
