package entity

import "gorm.io/gorm"

type Category struct {
	CoreEntity
	Name        string       `gorm:"type:varchar(255);not null" json:"name"`
	CoverImg    string       `gorm:"type:varchar(255);not null" json:"cover_img"`
	Restaurants []Restaurant `gorm:"foreignKey:CategoryID" json:"restaurants,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

// 유효성 검사 (옵션): 별도 함수로 구현 가능
func (c *Category) Validate() error {
	if len(c.Name) < 5 {
		return gorm.ErrInvalidData // "name must be at least 5 characters"
	}
	if c.Name == "" || c.CoverImg == "" {
		return gorm.ErrInvalidData // "name and coverImg are required"
	}
	return nil
}

// type CategoryResponse struct {
// 	CoreResponse
// 	Name        string               `json:"name,omitempty"`
// 	CoverImg    string               `json:"cover_img,omitempty"`
// 	Restaurants []RestaurantResponse `json:"restaurants,omitempty"`
// }

// func (c *Category) Serialize() CategoryResponse {
// 	restaurants := make([]RestaurantResponse, len(c.Restaurants))
// 	for i, restaurant := range c.Restaurants {
// 		restaurants[i] = restaurant.Serialize()
// 	}

// 	return CategoryResponse{
// 		Name:        c.Name,
// 		CoverImg:    c.CoverImg,
// 		Restaurants: restaurants,
// 	}
// }
