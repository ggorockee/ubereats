package repository

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(c *fiber.Ctx, inputParm *restaurantDto.CreateCategoryIn) (*entity.Category, error)
	GetAllCategory(c *fiber.Ctx) (*[]entity.Category, error)
	GetCategory(c *fiber.Ctx, id uint) (*entity.Category, error)
}

type categoryRepository struct {
	dbConn *gorm.DB
}

// CreateCategory implements CategoryRepository.
func (r *categoryRepository) CreateCategory(c *fiber.Ctx, inputParm *restaurantDto.CreateCategoryIn) (*entity.Category, error) {
	var category entity.Category

	if err := common.DecodeStructure(inputParm, &category); err != nil {
		return nil, err
	}

	if err := r.dbConn.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// GetAllCategory implements CategoryRepository.
func (r *categoryRepository) GetAllCategory(c *fiber.Ctx) (*[]entity.Category, error) {
	var categories []entity.Category

	if err := r.dbConn.Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

// GetCategory implements CategoryRepository.
func (r *categoryRepository) GetCategory(c *fiber.Ctx, id uint) (*entity.Category, error) {
	var category entity.Category

	if err := r.dbConn.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func NewCategoryRepository(
	dbConn *gorm.DB,
) CategoryRepository {
	return &categoryRepository{
		dbConn: dbConn,
	}
}
