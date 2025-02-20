package dto

type CreateDishInput struct {
	RestaurantID string       `json:"restaurantId" validate:"required"`
	Name         string       `json:"name" validate:"required,min=5"`
	Price        int          `json:"price" validate:"required"`
	Description  string       `json:"description" validate:"required,min=5,max=140"`
	Options      []DishOption `json:"options,omitempty"`
}

type EditDishInput struct {
	DishID      string       `json:"dishId" validate:"required"`
	Name        string       `json:"name,omitempty" validate:"min=5"`
	Price       int          `json:"price,omitempty"`
	Description string       `json:"description,omitempty" validate:"min=5,max=140"`
	Options     []DishOption `json:"options,omitempty"`
}

type DeleteDishInput struct {
	DishID string `json:"dishId" validate:"required"`
}

type DishChoice struct {
	Name  string `json:"name" validate:"required"`
	Extra int    `json:"extra,omitempty"`
}

type DishOption struct {
	Name    string       `json:"name" validate:"required"`
	Choices []DishChoice `json:"choices,omitempty"`
	Extra   int          `json:"extra,omitempty"`
}

type Dish struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Price        int          `json:"price"`
	Photo        string       `json:"photo,omitempty"`
	Description  string       `json:"description"`
	Options      []DishOption `json:"options,omitempty"`
	RestaurantID int          `json:"restaurantId"`
}
