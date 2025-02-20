package repository

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"

	categoryDto "ubereats/app/domain/category/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory(context ...*fiber.Ctx) (*[]entity.Category, error)
	CreateCategory(input *categoryDto.CreateCategory, c *fiber.Ctx) (*entity.Category, error)
	UpdateCategory(input *categoryDto.UpdateCategory, id int, c *fiber.Ctx) (*entity.Category, error)
	FindCategoryById(id int) (*entity.Category, error)
	DeleteCategory(id int, c *fiber.Ctx) error
	CountRestaurants(category *entity.Category) (int, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func (r *categoryRepository) CountRestaurants(category *entity.Category) (int, error) {
	var count int64
	if err := r.db.Model(&entity.Restaurant{}).Where("category_id = ?", category.ID).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *categoryRepository) DeleteCategory(id int, c *fiber.Ctx) error {
	category, err := r.GetFindById(id)
	if err != nil {
		return err
	}

	if err := r.db.Delete(category).Error; err != nil {
		return err
	}

	return nil
}

// GetFindById implements CategoryRepository.
func (r *categoryRepository) GetFindById(id int, context ...*fiber.Ctx) (*entity.Category, error) {
	var category entity.Category
	if err := r.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// UpdateCategory implements CategoryRepository.
func (r *categoryRepository) UpdateCategory(input *categoryDto.UpdateCategory, id int, c *fiber.Ctx) (*entity.Category, error) {
	category, err := r.GetFindById(id)
	if err != nil {
		return nil, err
	}

	if err := common.DecodeStructure(input, category); err != nil {
		return nil, err
	}

	if err := r.db.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

// CreateCategory implements CategoryRepository.
func (r *categoryRepository) CreateCategory(input *categoryDto.CreateCategory, c *fiber.Ctx) (*entity.Category, error) {
	var category entity.Category
	if err := common.DecodeStructure(input, &category); err != nil {
		return nil, err
	}

	if err := r.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// GetAll implements CategoryRepository.
func (r *categoryRepository) GetAllCategory(context ...*fiber.Ctx) (*[]entity.Category, error) {
	var categorys []entity.Category
	if err := r.db.Find(&categorys).Error; err != nil {
		return nil, err
	}
	return &categorys, nil
}

func NewCategoryRepository(d *gorm.DB) CategoryRepository {
	return &categoryRepository{db: d}
}
