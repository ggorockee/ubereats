package dto

// import (
// 	"ubereats/app/core/entity"
// )

// type CreateAccount struct {
// 	Email     string          `json:"email" validate:"email"`
// 	Password  string          `json:"password" validate:"required,min=8"`
// 	Password2 string          `json:"password2" validate:"required,eqfield=Password"`
// 	Role      entity.UserRole `json:"role" validate:"required,role" gorm:"default:'client'"`
// }

// // type UpdateUser struct {
// // 	Name       string `gorm:"column:name" json:"name,omitempty"`
// // 	IsVegan    bool   `gorm:"column:is_vegan" json:"is_vegan,omitempty"`
// // 	Address    string `gorm:"column:address" json:"address,omitempty"`
// // 	OwnersName string `gorm:"column:owners_name" json:"owners_name,omitempty"`
// // }
