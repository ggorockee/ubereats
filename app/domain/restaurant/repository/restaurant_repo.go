package repository

import (
	"ubereats/app/core/entity"

	restDto "ubereats/app/domain/restaurant/dto"

	"gorm.io/gorm"
)

type RestaurantRepository interface {
	CreateRestaurant(restaurant *restDto.Restaurant) (*entity.Restaurant, error)
	FindRestaurantByID(id int) (*entity.Restaurant, error)
	FindAll(page int) (*[]entity.Restaurant, *int, error)
	FindByName(query string, page int) (*[]entity.Restaurant, *int, error)
	UpdateRestaurant(id int, restaurant *restDto.Restaurant) (*entity.Restaurant, error)
	DeleteRestaurant(id int) error
	CreateDish(dish *restDto.Dish) (*entity.Dish, error)
	FindDishByID(id int) (*entity.Dish, error)
	UpdateDish(id int, dish *restDto.Dish) (*entity.Dish, error)
	DeleteDish(id int) error
}

type restaurantRepository struct {
	db *gorm.DB
}

// CreateDish implements RestaurantRepository.
func (r *restaurantRepository) CreateDish(dish *restDto.Dish) (*entity.Dish, error) {
	entityDish := entity.Dish{
		Name:         dish.Name,
		Price:        dish.Price,
		Photo:        dish.Photo,
		Description:  dish.Description,
		RestaurantID: dish.RestaurantID,
	}

	if err := r.db.Create(&entityDish).Error; err != nil {
		return nil, err
	}

	return &entityDish, nil
}

// CreateRestaurant implements RestaurantRepository.
func (r *restaurantRepository) CreateRestaurant(restaurant *restDto.Restaurant) (*entity.Restaurant, error) {
	entityRestaurant := entity.Restaurant{
		Name:     restaurant.Name,
		CoverImg: restaurant.CoverImg,
		Address:  restaurant.Address,
	}

	// (TODO) 관계 설정

	if err := r.db.Create(&entityRestaurant).Error; err != nil {
		return nil, err
	}

	return &entityRestaurant, nil
}

// DeleteDish implements RestaurantRepository.
func (r *restaurantRepository) DeleteDish(id int) error {
	dish, err := r.FindDishByID(id)
	if err != nil {
		return err
	}
	if err := r.db.Delete(dish).Error; err != nil {
		return err
	}
	return nil
}

// DeleteRestaurant implements RestaurantRepository.
func (r *restaurantRepository) DeleteRestaurant(id int) error {
	restaurant, err := r.FindRestaurantByID(id)
	if err != nil {
		return err
	}
	if err := r.db.Delete(restaurant).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements RestaurantRepository.
func (r *restaurantRepository) FindAll(page int) (*[]entity.Restaurant, *int, error) {
	var restaurants []entity.Restaurant
	var total int64

	if err := r.db.Model(&entity.Restaurant{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	if err := r.db.Offset((page - 1) * 25).Limit(25).Find(&restaurants).Error; err != nil {
		return nil, nil, err
	}

	t := int(total)
	return &restaurants, &t, nil
}

// FindByName implements RestaurantRepository.
func (r *restaurantRepository) FindByName(query string, page int) (*[]entity.Restaurant, *int, error) {
	var restaurants []entity.Restaurant
	var total int64
	if err := r.db.Model(&restDto.Restaurant{}).Where("name LIKE ?", "%"+query+"%").Count(&total).Error; err != nil {
		return nil, nil, err
	}
	if err := r.db.Where("name LIKE ?", "%"+query+"%").Offset((page - 1) * 25).Limit(25).Find(&restaurants).Error; err != nil {
		return nil, nil, err
	}

	t := int(total)
	return &restaurants, &t, nil
}

// FindDishByID implements RestaurantRepository.
func (r *restaurantRepository) FindDishByID(id int) (*entity.Dish, error) {
	var dish entity.Dish
	if err := r.db.Preload("Restaurant").
		Where("id = ?", id).
		First(&dish).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

// FindRestaurantByID implements RestaurantRepository.
func (r *restaurantRepository) FindRestaurantByID(id int) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant
	if err := r.db.Preload("Menu").
		Where("id = ?", id).
		First(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

// UpdateDish implements RestaurantRepository.
func (r *restaurantRepository) UpdateDish(id int, dish *restDto.Dish) (*entity.Dish, error) {
	entityDish, err := r.FindDishByID(id)
	if err != nil {
		return nil, err
	}
	if err := r.db.Model(entityDish).Updates(dish).Error; err != nil {
		return nil, err
	}

	return entityDish, nil
}

// UpdateRestaurant implements RestaurantRepository.
func (r *restaurantRepository) UpdateRestaurant(id int, restaurant *restDto.Restaurant) (*entity.Restaurant, error) {
	entityRestaurant, err := r.FindRestaurantByID(id)
	if err != nil {
		return nil, err
	}
	if err := r.db.Model(entityRestaurant).Updates(restaurant).Error; err != nil {
		return nil, err
	}

	return entityRestaurant, nil
}

func NewRestaurantRepository(d *gorm.DB) RestaurantRepository {
	return &restaurantRepository{
		db: d,
	}
}
