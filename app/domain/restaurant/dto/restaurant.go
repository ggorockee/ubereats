package dto

type CreateRestaurantDto struct {
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Category string `json:"category" validate:"required"`
}
