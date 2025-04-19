package entity

// User는 사용자 엔티티를 나타냅니다.
type User struct {
	CoreEntity          // CoreEntity 임베딩
	Email      string   `gorm:"type:varchar(255);unique;not null" json:"email" validate:"email" mapstructure:"email"`
	Password   string   `gorm:"type:varchar(255);not null" json:"-" mapstructure:"password"`
	Role       UserRole `gorm:"type:varchar(20);not null;default:'client'" json:"role" validate:"role" mapstructure:"role"`
	// Restaurants []Restaurant `gorm:"foreignKey:OwnerID" json:"restaurants"` // 1:N 관계

	Restaurants []Restaurant `gorm:"foreignKey:OwnerRefer" json:"restaurants,omitempty" mapstructure:"restaurants"`
}

func (m *User) UpdateDelProperty() {
	m.IsDel = "Y"
}

func (User) TableName() string {
	return "user"
}
