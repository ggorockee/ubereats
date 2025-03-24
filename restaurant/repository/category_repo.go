package repository

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(input *restaurantDto.CreateCategoryInput, c *fiber.Ctx) (*entity.Category, error)
	FindById(id string) (*entity.Category, error)
	FindAll() (*[]entity.Category, error)
	Update(input *restaurantDto.EditCategoryInput, c *fiber.Ctx) (*entity.Category, error)
	Delete(id string) error
}

type categoryRepository struct {
	dbConn *gorm.DB
}

// Create implements CategoryRepository.
func (r *categoryRepository) Create(input *restaurantDto.CreateCategoryInput, c *fiber.Ctx) (*entity.Category, error) {
	var category entity.Category
	if err := common.DecodeStructure(input, &category); err != nil {
		return nil, err
	}

	err := r.dbConn.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Delete implements CategoryRepository.
func (r *categoryRepository) Delete(id string) error {
	category, err := r.FindById(id)
	if err != nil {
		return err
	}

	err = r.dbConn.Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements CategoryRepository.
func (r *categoryRepository) FindAll() (*[]entity.Category, error) {
	var categories []entity.Category
	err := r.dbConn.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

// GetFindById implements CategoryRepository.
func (r *categoryRepository) FindById(id string) (*entity.Category, error) {
	var category entity.Category

	err := r.dbConn.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Update implements CategoryRepository.
func (r *categoryRepository) Update(input *restaurantDto.EditCategoryInput, c *fiber.Ctx) (*entity.Category, error) {
	category, err := r.FindById(input.CategoryID)
	if err != nil {
		return nil, err
	}

	if err := common.DecodeStructure(input, category); err != nil {
		return nil, err
	}

	err = r.dbConn.Model(&entity.Category{}).Save(category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}

func NewCategoryRepository(d *gorm.DB) CategoryRepository {
	return &categoryRepository{
		dbConn: d,
	}
}
