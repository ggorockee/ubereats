package dto

type CreateRestaurantIn struct {
	Name       string `json:"name" validate:"required,min=2" mapstructure:"name"`
	Address    string `json:"address" validate:"required" mapstructure:"address"`
	CoverImg   string `json:"cover_img" validate:"required" mapstructure:"cover_img"`
	CategoryId uint   `json:"category_id" validate:"required" mapstructure:"category_id"`
}

type EditRestaurantIn struct {
	Name       *string `json:"name,omitempty" validate:"required,min=2" mapstructure:"name"`
	Address    *string `json:"address,omitempty" validate:"required" mapstructure:"address"`
	CoverImg   *string `json:"cover_img,omitempty" validate:"required" mapstructure:"cover_img"`
	CategoryId *uint   `json:"category_id,omitempty" validate:"required" mapstructure:"category_id"`
}

type DeleteRestaurantInput struct {
	RestaurantId uint `json:"restaurant_id"`
}
