package entity

type UserRole string

func (r UserRole) IsValid() bool {
	switch r {
	case RoleClient, RoleOwner, RoleDelivery:
		return true
	default:
		return false
	}
}

// UserRole 상수 정의
const (
	RoleClient   UserRole = "Client"
	RoleOwner    UserRole = "Owner"
	RoleDelivery UserRole = "Delivery"
	RoleAny      UserRole = "Any"
)

// User는 사용자 엔티티를 나타냅니다.
type User struct {
	CoreEntity               // CoreEntity 임베딩
	Email       string       `gorm:"type:varchar(255);unique;not null" json:"email" validate:"email"`
	Password    string       `gorm:"type:varchar(255);not null" json:"-"`
	Role        UserRole     `gorm:"type:varchar(20);not null" json:"role"`
	Restaurants []Restaurant `gorm:"foreignKey:OwnerID" json:"restaurants"` // 1:N 관계
}

func (User) TableName() string {
	return "users"
}
