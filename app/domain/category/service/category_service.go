package service

import (
	"ubereats/app/core/entity"

	categoryDto "ubereats/app/domain/category/dto"
	categoryRepo "ubereats/app/domain/category/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryService interface {
	GetAllCategory(context ...*fiber.Ctx) (*[]entity.Category, error)
	CreateCategory(input *categoryDto.CreateCategory, c *fiber.Ctx) (*entity.Category, error)
	UpdateCategory(input *categoryDto.UpdateCategory, id int, c *fiber.Ctx) (*entity.Category, error)
	GetFindById(id int, context ...*fiber.Ctx) (*entity.Category, error)
	DeleteCategory(id int, c *fiber.Ctx) error
	CountRestaurants(category *entity.Category) (int, error)
}

type categoryService struct {
	db           *gorm.DB
	categoryRepo categoryRepo.CategoryRepository
}

func (s *categoryService) CountRestaurants(category *entity.Category) (int, error) {
	var count int
	var err error
	err = s.db.Transaction(func(tx *gorm.DB) error {
		count, err = s.categoryRepo.CountRestaurants(category)
		if err != nil {
			return err
		}
		return nil
	})
	return count, err
}

// DeleteCategory implements CategoryService.
func (s *categoryService) DeleteCategory(id int, c *fiber.Ctx) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.categoryRepo.DeleteCategory(id, c)
	})

	return err
}

// GetFindById implements CategoryService.
func (s *categoryService) GetFindById(id int, context ...*fiber.Ctx) (*entity.Category, error) {
	var err error
	var category *entity.Category

	err = s.db.Transaction(func(tx *gorm.DB) error {
		category, err = s.categoryRepo.GetFindById(id, context...)
		if err != nil {
			return err
		}
		return nil
	})

	return category, err
}

// UpdateCategory implements CategoryService.
func (s *categoryService) UpdateCategory(input *categoryDto.UpdateCategory, id int, c *fiber.Ctx) (*entity.Category, error) {
	var err error
	var category *entity.Category
	err = s.db.Transaction(func(tx *gorm.DB) error {
		category, err = s.categoryRepo.UpdateCategory(input, id, c)
		if err != nil {
			return err
		}
		return nil
	})
	return category, err

}

// CreateCategory implements CategoryService.
func (s *categoryService) CreateCategory(input *categoryDto.CreateCategory, c *fiber.Ctx) (*entity.Category, error) {
	var err error
	var category *entity.Category
	err = s.db.Transaction(func(tx *gorm.DB) error {
		category, err = s.categoryRepo.CreateCategory(input, c)
		if err != nil {
			return err
		}
		return nil
	})
	return category, err
}

// GetAll implements CategoryService.
func (s *categoryService) GetAllCategory(context ...*fiber.Ctx) (*[]entity.Category, error) {
	var err error
	var categorys *[]entity.Category

	err = s.db.Transaction(func(tx *gorm.DB) error {
		categorys, err = s.categoryRepo.GetAllCategory(context...)
		if err != nil {
			return err
		}
		return nil
	})

	return categorys, err
}

func NewCategoryService(d *gorm.DB, r categoryRepo.CategoryRepository) CategoryService {
	return &categoryService{
		db:           d,
		categoryRepo: r,
	}
}
