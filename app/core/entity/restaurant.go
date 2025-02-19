package entity

import "gorm.io/gorm"

type Restaurant struct {
	CoreEntity
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	CoverImg   string    `gorm:"type:varchar(255);not null" json:"cover_img"`
	Address    string    `gorm:"type:varchar(255);not null;default:'강남'" json:"address"`
	CategoryID *int      `gorm:"index" json:"-"`                                  // 외래 키 (nullable)
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"` // @ManyToOne, nullable, onDelete: SET NULL
	OwnerID    int       `gorm:"not null" json:"-"`                               // 외래 키 (User와 연결)
	Owner      User      `gorm:"foreignKey:OwnerID" json:"owner"`                 // @ManyToOne
}

func (Restaurant) TableName() string {
	return "restaurants"
}

// 유효성 검사 (옵션)
func (r *Restaurant) Validate() error {
	if len(r.Name) < 5 {
		return gorm.ErrInvalidData // "name must be at least 5 characters"
	}
	if r.Name == "" || r.CoverImg == "" || r.Address == "" {
		return gorm.ErrInvalidData // "name, coverImg, and address are required"
	}
	if r.OwnerID == 0 {
		return gorm.ErrInvalidData // "owner is required"
	}
	return nil
}
