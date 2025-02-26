package entity

import (
	"errors"
)

type Restaurant struct {
	CoreEntity
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	CoverImg   string    `gorm:"type:varchar(255);not null" json:"cover_img"`
	Address    string    `gorm:"type:varchar(255);not null;default:'강남'" json:"address"`
	CategoryID int       `gorm:"index" json:"category_id"`                        // 외래 키 (nullable)
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"` // @ManyToOne, nullable, onDelete: SET NULL
	OwnerID    int       `gorm:"not null" json:"owner_id"`                        // 외래 키 (User와 연결)
	Owner      User      `gorm:"foreignKey:OwnerID" json:"owner"`                 // @ManyToOne
	Menu       []Dish    `gorm:"foreignKey:RestaurantID" json:"menu,omitempty"`
}

type RestaurantResponse struct {
	CoreResponse
	Name     string            `json:"name"`
	CoverImg string            `json:"cover_img"`
	Address  string            `json:"address"`
	Category *CategoryResponse `json:"category,omitempty"` // @ManyToOne, nullab
	Owner    UserResponse      `json:"owner"`              // @ManyToOne
	Menu     []DishResponse    `json:"menu,omitempty"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

// 유효성 검사 (옵션)
func (r *Restaurant) Validate() error {

	switch {
	case r.Name == "":
		return errors.New("name is required")
	case r.CoverImg == "":
		return errors.New("cover_img is required")
	case r.Address == "":
		return errors.New("address is required")
	case r.OwnerID == 0:
		return errors.New("owner_id is required")
	case r.CategoryID == 0:
		return errors.New("category_id must be a positive integer if provided")
	}

	return nil
}

func (r *Restaurant) Serialize() RestaurantResponse {
	categoryResponse := r.Category.Serialize()
	user := r.Owner.Serialize()
	return RestaurantResponse{
		Name:     r.Name,
		CoverImg: r.CoverImg,
		Address:  r.Address,
		Category: &categoryResponse,
		Owner:    user,
		Menu:     r.Menu,
	}
}
