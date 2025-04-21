package dto

type CreateRestaurantIn struct {
	Name       string `json:"name" validate:"required,min=2" mapstructure:"name"`
	Address    string `json:"address" validate:"required" mapstructure:"address"`
	CoverImg   string `json:"cover_img" validate:"required" mapstructure:"cover_img"`
	CategoryId uint   `json:"category_id" validate:"required" mapstructure:"category_id"`
}

type EditRestaurantIn struct {
	Name       *string `json:"name,omitempty" validate:"min=2" mapstructure:"name"`
	Address    *string `json:"address,omitempty" mapstructure:"address"`
	CoverImg   *string `json:"cover_img,omitempty"  mapstructure:"cover_img"`
	CategoryId *uint   `json:"category_id,omitempty"  mapstructure:"category_id"`
}

type DeleteRestaurantInput struct {
	RestaurantId uint `json:"restaurant_id"`
}

type GetRestaurantsInput struct {
	Page  int `json:"page" query:"page" validate:"min=1"`
	Limit int `json:"limit" query:"limit" validate:"min=1"`
}
