package repository

import (
	"errors"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"

	categoryRepo "ubereats/app/domain/category/repository"
	dishDto "ubereats/app/domain/dish/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DishRepository interface {
	CreateDish(input *dishDto.CreateDish, c *fiber.Ctx) (*entity.Dish, error)
	UpdateDish(input *dishDto.UpdateDish, id int, c *fiber.Ctx) (*entity.Dish, error)
	DeleteDish(id int, c *fiber.Ctx) error
	AllDishs(input dishDto.DishsInput) (*[]entity.Dish, error)
	FindDishById(input dishDto.DishInput) (*entity.Dish, error)
	SearchDishByName(input dishDto.SearchDish) (*[]entity.Dish, error)
	AllCategories() (*[]entity.Category, error)
	CountDishs(category entity.Category) (int, error)
}

type dishRepository struct {
	db           *gorm.DB
	categoryRepo categoryRepo.CategoryRepository
}

func (r *dishRepository) SearchDishByName(input dishDto.SearchDish) (*[]entity.Dish, error) {
	var dishs []entity.Dish
	if err := r.db.
		Where("name LIKE ?", "%"+input.Query+"%").
		Limit(25).
		Offset((input.Page - 1) * 25).
		Find(&dishs).Error; err != nil {
		return nil, err
	}

	return &dishs, nil
}

func (r *dishRepository) FindDishById(input dishDto.DishInput) (*entity.Dish, error) {
	var dish entity.Dish
	if err := r.db.First(&dish, input.DishID).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

func (r *dishRepository) CountDishs(category entity.Category) (int, error) {
	var count int64
	if err := r.db.
		Model(&entity.Dish{}).
		Where("category_id = ?", category.ID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *dishRepository) AllDishs(input dishDto.DishsInput) (*[]entity.Dish, error) {
	var dishs []entity.Dish
	if err := r.db.
		Limit(25).
		Offset((input.Page - 1) * 25).
		Preload("Owner").
		Preload("Category").
		Find(&dishs).Error; err != nil {
		return nil, errors.New("could not load dishs")
	}

	return &dishs, nil
}

func (r *dishRepository) AllCategories() (*[]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *dishRepository) DeleteDish(id int, c *fiber.Ctx) error {
	dish, err := r.GetFindById(id, c)
	if err != nil {
		return err
	}

	if err := r.db.Delete(dish).Error; err != nil {
		return err
	}

	return nil
}

// GetFindById implements DishRepository.
func (r *dishRepository) GetFindById(id int, context ...*fiber.Ctx) (*entity.Dish, error) {
	var dish entity.Dish
	if err := r.db.Preload("Owner").Preload("Category").Where("id = ?", id).First(&dish).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

// UpdateDish implements DishRepository.
func (r *dishRepository) UpdateDish(input *dishDto.UpdateDish, id int, c *fiber.Ctx) (*entity.Dish, error) {
	dish, err := r.GetFindById(id, c)
	if err != nil {
		return nil, err
	}

	if err := common.DecodeStructure(input, dish); err != nil {
		return nil, err
	}

	if err := dish.Validate(); err != nil {
		return nil, err
	}

	category, err := r.categoryRepo.GetFindById(dish.CategoryID)
	if err != nil {
		return nil, err
	}

	dish.Category = category

	if err := r.db.Save(dish).Error; err != nil {
		return nil, err
	}

	if err := r.db.Preload("Owner").Preload("Category").Where("id = ?", id).First(dish).Error; err != nil {
		return nil, err
	}

	return dish, nil
}

// CreateDish implements DishRepository.
func (r *dishRepository) CreateDish(input *dishDto.CreateDish, c *fiber.Ctx) (*entity.Dish, error) {

	user, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, errors.New("user not authenticated")
	}

	var dish entity.Dish

	if err := common.DecodeStructure(input, &dish); err != nil {
		return nil, err
	}

	dish.OwnerID = user.ID

	if err := dish.Validate(); err != nil {
		return nil, err
	}

	if err := r.db.Create(&dish).Error; err != nil {
		return nil, err
	}

	if err := r.db.
		Preload("Category").
		Preload("Owner").
		First(&dish).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

func NewDishRepository(d *gorm.DB, c categoryRepo.CategoryRepository) DishRepository {
	return &dishRepository{
		db:           d,
		categoryRepo: c,
	}
}
