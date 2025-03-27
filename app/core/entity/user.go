package entity

// func (r UserRole) IsValid() bool {
// 	switch r {
// 	case RoleClient, RoleOwner, RoleDelivery:
// 		return true
// 	default:
// 		return false
// 	}
// }

// // // UserRole 상수 정의
// const (
// 	RoleClient   UserRole = "Client"
// 	RoleOwner    UserRole = "Owner"
// 	RoleDelivery UserRole = "Delivery"
// 	RoleAny      UserRole = "Any"
// )

// User는 사용자 엔티티를 나타냅니다.
type User struct {
	CoreEntity        // CoreEntity 임베딩
	Email      string `gorm:"type:varchar(255);unique;not null" json:"email" validate:"email"`
	Password   string `gorm:"type:varchar(255);not null" json:"-"`
	Role       string `gorm:"type:varchar(20);not null" json:"role" validate:"role"`
	// Restaurants []Restaurant `gorm:"foreignKey:OwnerID" json:"restaurants"` // 1:N 관계
}

// type UserResponse struct {
// 	CoreResponse
// 	Email       string               `json:"email"`
// 	Role        UserRole             `json:"role"`
// 	Restaurants []RestaurantResponse `json:"restaurants"` // 1:N 관계
// }

// func (u *User) Serialize() UserResponse {
// 	restaurants := make([]RestaurantResponse, len(u.Restaurants))

// 	for i, resrestaurant := range u.Restaurants {
// 		restaurants[i] = resrestaurant.Serialize()
// 	}

// 	return UserResponse{
// 		Email:       u.Email,
// 		Role:        u.Role,
// 		Restaurants: restaurants,
// 	}
// }

// func (User) TableName() string {
// 	return "users"
// }
