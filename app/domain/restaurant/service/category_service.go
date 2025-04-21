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

type CategoryService interface {
	CreateCategory(c *fiber.Ctx, inputParm *restaurantDto.CreateCategoryIn) (*restaurantResp.CreateCategoryOut, error)
	GetAllCategory(c *fiber.Ctx) (*restaurantResp.GetAllCategoryOut, error)
	GetCategory(c *fiber.Ctx, id uint) (*restaurantResp.GetCategoryOut, error)
}

type categoryService struct {
	categoryRepo restaurantRepo.CategoryRepository
	dbConn       *gorm.DB
}

// CreateCategory implements CategoryService.
func (s *categoryService) CreateCategory(c *fiber.Ctx, inputParm *restaurantDto.CreateCategoryIn) (*restaurantResp.CreateCategoryOut, error) {
	var category *entity.Category
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		category, err = s.categoryRepo.CreateCategory(c, inputParm)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantResp.CreateCategoryOut{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &restaurantResp.CreateCategoryOut{
		CoreResponse: common.CoreResponse{
			Ok:   true,
			Data: category,
		},
	}, nil
}

// GetAllCategory implements CategoryService.
func (s *categoryService) GetAllCategory(c *fiber.Ctx) (*restaurantResp.GetAllCategoryOut, error) {
	categories, err := s.categoryRepo.GetAllCategory(c)
	if err != nil {
		return &restaurantResp.GetAllCategoryOut{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &restaurantResp.GetAllCategoryOut{
		CoreResponse: common.CoreResponse{
			Ok:   true,
			Data: categories,
		},
	}, nil
}

// GetCategory implements CategoryService.
func (s *categoryService) GetCategory(c *fiber.Ctx, id uint) (*restaurantResp.GetCategoryOut, error) {
	category, err := s.categoryRepo.GetCategory(c, id)
	if err != nil {
		return &restaurantResp.GetCategoryOut{
			Message: err.Error(),
		}, err
	}

	// 레스토랑 목록 변환
	simpleRestaurants := restaurantResp.ToSimpleRestaurants(category.Restaurants)

	// 카테고리 결과 변환
	result := &restaurantResp.CategoryResult{
		ID:       category.ID,
		Name:     category.Name,
		CoverImg: category.CoverImg,
		// 레스토랑 목록은 필요시 변환
		Restaurants: simpleRestaurants,
	}

	return &restaurantResp.GetCategoryOut{
		Ok:   true,
		Data: result,
	}, nil
}

func NewCategoryService(
	categoryRepo restaurantRepo.CategoryRepository,
	dbConn *gorm.DB,
) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
		dbConn:       dbConn,
	}
}
