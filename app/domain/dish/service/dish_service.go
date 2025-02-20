package service

import (
	"fmt"
	"ubereats/app/core/entity"

	dishDto "ubereats/app/domain/dish/dto"
	dishRepo "ubereats/app/domain/dish/repository"
	dishRes "ubereats/app/domain/dish/response"

	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DishService interface {
	CreateDish(input *dishDto.CreateDish, c *fiber.Ctx) (*entity.Dish, error)
	UpdateDish(input *dishDto.UpdateDish, id int, c *fiber.Ctx) (*entity.Dish, error)
	DeleteDish(id int, c *fiber.Ctx) error
	AllDishs(input dishDto.DishsInput) (*dishRes.DishsResponse, error)
	FindDishById(input dishDto.DishInput) (*dishRes.DishResponse, error)
	SearchDishByName(input dishDto.SearchDish) (*dishRes.SearchDishResponse, error)
	AllCategories() (*dishRes.AllCategoriesResponse, error)
	CountDishs(category entity.Category) (int, error)
}

type dishService struct {
	db             *gorm.DB
	dishRepo       dishRepo.DishRepository
	restaurantRepo restaurantRepo.RestaurantRepository
}

// AllCategories implements DishService.
func (s *dishService) AllCategories() (*dishRes.AllCategoriesResponse, error) {
	var categories *[]entity.Category

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		categories, err = s.dishRepo.AllCategories()
		if err != nil {
			return err
		}
		return nil
	})

	return &dishRes.AllCategoriesResponse{
		BaseResponse: dishRes.BaseResponse{Ok: true},
		Categories:   *categories,
	}, err
}

// AllDishs implements DishService.
func (s *dishService) AllDishs(input dishDto.DishsInput) (*dishRes.DishsResponse, error) {
	var dishs *[]entity.Dish
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		dishs, err = s.dishRepo.AllDishs(input)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &dishRes.DishsResponse{
			BaseResponse: dishRes.BaseResponse{
				Ok:    false,
				Error: "Could not load dishs",
			},
		}, err
	}

	var totalResults int64
	err = s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Dish{}).Count(&totalResults).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &dishRes.DishsResponse{
			BaseResponse: dishRes.BaseResponse{
				Ok:    false,
				Error: "Could not load dishs",
			},
		}, err
	}

	return &dishRes.DishsResponse{
		BaseResponse: dishRes.BaseResponse{Ok: true},
		Results:      *dishs,
		TotalPages:   int(totalResults+24) / 25, // ceil 처리
		TotalResults: int(totalResults),
	}, err
}

// CountDishs implements DishService.
func (s *dishService) CountDishs(category entity.Category) (int, error) {
	var count int
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		count, err = s.dishRepo.CountDishs(category)
		if err != nil {
			return err
		}
		return nil
	})
	return count, err
}

// CreateDish implements DishService.
func (s *dishService) CreateDish(input *dishDto.CreateDish, c *fiber.Ctx) (*entity.Dish, error) {
	user, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}

	inputDto := restaurantDto.RestaurantInput{RestaurantID: input.RestaurantID}
	restaurant, err := s.restaurantRepo.FindRestaurantById(inputDto)
	if err != nil {
		return nil, err
	}

	if restaurant.OwnerID != user.ID {
		return nil, fmt.Errorf("you can't create a dish for a restaurant you don't own")
	}

	var dish *entity.Dish
	err = s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		dish, err = s.dishRepo.CreateDish(input, c)
		return err
	})
	return dish, err
}

// DeleteDish implements DishService.
func (s *dishService) DeleteDish(id int, c *fiber.Ctx) error {
	user, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return fmt.Errorf("user not authenticated")
	}

	dish, err := s.dishRepo.FindDishById(dishDto.DishInput{DishID: id})
	if err != nil {
		return err
	}
	if dish.Restaurant.OwnerID != user.ID {
		return fmt.Errorf("you can't delete a dish for a restaurant you don't own")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.dishRepo.DeleteDish(id, c)
	})
}

// FindDishById implements DishService.
func (s *dishService) FindDishById(input dishDto.DishInput) (*dishRes.DishResponse, error) {
	var dish *entity.Dish
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		dish, err = s.dishRepo.FindDishById(input)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &dishRes.DishResponse{
			BaseResponse: dishRes.BaseResponse{
				Ok:    false,
				Error: "Dish not found",
			},
		}, err
	}

	return &dishRes.DishResponse{
		BaseResponse: dishRes.BaseResponse{Ok: true},
		Dish:         dish,
	}, nil
}

// SearchDishByName implements DishService.
func (s *dishService) SearchDishByName(input dishDto.SearchDish) (*dishRes.SearchDishResponse, error) {
	var dishs *[]entity.Dish
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		dishs, err = s.dishRepo.SearchDishByName(input)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &dishRes.SearchDishResponse{
			BaseResponse: dishRes.BaseResponse{Ok: true, Error: "Could not search for dishs"},
		}, err
	}

	var totalResults int64
	err = s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Dish{}).Where("name LIKE ?", "%"+input.Query+"%").Count(&totalResults).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &dishRes.SearchDishResponse{
			BaseResponse: dishRes.BaseResponse{Ok: false, Error: "Could not count dishs"},
			Dishs:        *dishs,
			TotalPages:   0,
			TotalResults: 0,
		}, err
	}

	return &dishRes.SearchDishResponse{
		BaseResponse: dishRes.BaseResponse{Ok: true},
		Dishs:        *dishs,
		TotalPages:   int(totalResults+24) / 25,
		TotalResults: int(totalResults),
	}, nil
}

// UpdateDish implements DishService.
func (s *dishService) UpdateDish(input *dishDto.UpdateDish, id int, c *fiber.Ctx) (*entity.Dish, error) {

	user, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}

	var dish *entity.Dish
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		dish, err = s.dishRepo.UpdateDish(input, id, c)
		if err != nil {
			return err
		}
		if dish.Restaurant.OwnerID != user.ID {
			return fmt.Errorf("you can't update a dish for a restaurant you don't own")
		}
		return nil
	})
	return dish, err
}

func NewDishService(d *gorm.DB, r dishRepo.DishRepository, restaurantRepo restaurantRepo.RestaurantRepository) DishService {
	return &dishService{
		db:             d,
		dishRepo:       r,
		restaurantRepo: restaurantRepo,
	}
}
