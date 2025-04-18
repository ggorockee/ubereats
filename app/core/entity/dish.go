package entity

type DishChoice struct {
	Name  string `json:"name"`
	Extra *int   `json:"extra,omitempty"`
}

type DishOption struct {
	Name    string       `json:"name"`
	Choices []DishChoice `json:"choices,omitempty"`
	Extra   *int         `json:"extra,omitempty"`
}

type Dish struct {
	CoreEntity
	Name         string       `gorm:"type:varchar(255);not null" json:"name"`
	Price        int          `gorm:"not null" json:"price"`
	Photo        string       `gorm:"type:varchar(255)" json:"photo,omitempty"`
	Description  string       `gorm:"type:varchar(255);not null" json:"description"`
	RestaurantID int          `gorm:"index" json:"restaurant_id"`
	Restaurant   Restaurant   `gorm:"foreignKey:RestaurantID" json:"restaurant,omitempty"`
	Options      []DishOption `gorm:"type:json" json:"options,omitempty"`
}

// type DishResponse struct {
// 	CoreResponse
// 	Name        string              `json:"name"`
// 	Price       int                 `json:"price"`
// 	Photo       string              `json:"photo,omitempty"`
// 	Description string              `json:"description"`
// 	Restaurant  *RestaurantResponse `json:"restaurant,omitempty"`
// 	Options     []DishOption        `json:"options,omitempty"`
// }

// func (d *Dish) Serialize() DishResponse {
// 	restaurantResponse := d.Restaurant.Serialize()
// 	return DishResponse{
// 		Name:        d.Name,
// 		Price:       d.Price,
// 		Photo:       d.Photo,
// 		Description: d.Description,
// 		Restaurant:  &restaurantResponse,
// 		Options:     d.Options,
// 	}
// }

// func (Dish) TableName() string {
// 	return "dishes"
// }

// func (d *Dish) Validate() error {
// 	if len(d.Name) < 5 || len(d.Description) < 5 || len(d.Description) > 140 {
// 		return errors.New("name must be at least 5 characters, description must be 5-140 characters")
// 	}
// 	return nil
// }
