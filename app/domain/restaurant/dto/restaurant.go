package dto

type CreateRestaurantInput struct {
	Name         string `json:"name" validate:"required,min=5"`
	CoverImg     string `json:"coverImg" validate:"required"`
	Address      string `json:"address" validate:"required"`
	CategoryName string `json:"categoryName" validate:"required"`
}

type EditRestaurantInput struct {
	RestaurantID string `json:"restaurantId" validate:"required"`
	Name         string `json:"name,omitempty" validate:"min=5"`
	CoverImg     string `json:"coverImg,omitempty"`
	Address      string `json:"address,omitempty"`
	CategoryName string `json:"categoryName,omitempty"`
}

type DeleteRestaurantInput struct {
	RestaurantID string `json:"restaurantId" validate:"required"`
}

type RestaurantsInput struct {
	Page int `query:"page" validate:"required,min=1"`
}

type RestaurantInput struct {
	RestaurantID string `json:"restaurantId" validate:"required"`
}

type SearchRestaurantInput struct {
	Query string `query:"query" validate:"required"`
	Page  int    `query:"page" validate:"required,min=1"`
}

type Restaurant struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	CoverImg string   `json:"coverImg"`
	Address  string   `json:"address"`
	Category Category `json:"category"`
	OwnerID  uint     `json:"ownerId"`
	Menu     []Dish   `json:"menu"`
}
