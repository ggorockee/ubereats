package entity

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
	Orders     []Order
}
