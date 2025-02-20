package service

import (
	"ubereats/app/core/entity"

	categoryDto "ubereats/app/domain/category/dto"
	categoryRepo "ubereats/app/domain/category/repository"
	categoryRes "ubereats/app/domain/category/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryService interface {
	GetAllCategory(c *fiber.Ctx) (*categoryRes.AllCategoriesResponse, error)
	CreateCategory(input *categoryDto.CreateCategory, c *fiber.Ctx) (*entity.Category, error)
	UpdateCategory(input *categoryDto.UpdateCategory, id int, c *fiber.Ctx) (*entity.Category, error)
	FindCategoryById(id int, c *fiber.Ctx) (*entity.Category, error)
	DeleteCategory(id int, c *fiber.Ctx) error
	CountRestaurants(category *entity.Category) (int, error)
}

type categoryService struct {
	db           *gorm.DB
	categoryRepo categoryRepo.CategoryRepository
}

// CountRestaurants implements CategoryService.
func (s *categoryService) CountRestaurants(category *entity.Category) (int, error) {
	panic("unimplemented")
}

// CreateCategory implements CategoryService.
func (s *categoryService) CreateCategory(input *categoryDto.CreateCategory, c *fiber.Ctx) (*categoryRes.CategoryResponse, error) {
	panic("unimplemented")
}

// DeleteCategory implements CategoryService.
func (s *categoryService) DeleteCategory(id int, c *fiber.Ctx) error {
	panic("unimplemented")
}

// FindCategoryById implements CategoryService.
func (s *categoryService) FindCategoryById(id int, c *fiber.Ctx) (*categoryRes.CategoryResponse, error) {
	var category *entity.Category
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		category, err = s.categoryRepo.FindCategoryById(id)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &categoryRes.CategoryResponse{
			BaseResponse: categoryRes.BaseResponse{Ok: false, Error: err.Error()},
		}, err
	}

	return &categoryRes.CategoryResponse{
		BaseResponse: categoryRes.BaseResponse{Ok: true},
		ID:           category.ID,
		CreatedAt:    "",
		UpdatedAt:    "",
		Name:         "",
		CoverImg:     "",
		Restaurants:  []entity.Restaurant{},
	}, nil
}

// GetAllCategory implements CategoryService.
func (s *categoryService) GetAllCategory(c *fiber.Ctx) (*categoryRes.AllCategoriesResponse, error) {
	var categories *[]entity.Category
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		categories, err = s.categoryRepo.GetAllCategory(c)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &categoryRes.AllCategoriesResponse{
			BaseResponse: categoryRes.BaseResponse{Ok: false, Error: err.Error()},
		}, err
	}

	counts := make([]int, len(*categories))
	for i, category := range *categories {
		count, err := s.categoryRepo.CountRestaurants(&category)
		if err != nil {
			return &categoryRes.AllCategoriesResponse{
				BaseResponse: categoryRes.BaseResponse{
					Ok:    false,
					Error: "Could not count restaurants",
				},
			}, err
		}
		counts[i] = count
	}

	result := categoryRes.GenCategoriesRes(categories, counts)
	return &categoryRes.AllCategoriesResponse{
		BaseResponse: categoryRes.BaseResponse{Ok: true},
		Categories:   result,
	}, nil
}

// UpdateCategory implements CategoryService.
func (s *categoryService) UpdateCategory(input *categoryDto.UpdateCategory, id int, c *fiber.Ctx) (*entity.Category, error) {
	panic("unimplemented")
}

func NewCategoryService(d *gorm.DB, r categoryRepo.CategoryRepository) CategoryService {
	return &categoryService{
		db:           d,
		categoryRepo: r,
	}
}
