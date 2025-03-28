package dto

import "ubereats/app/core/entity"

type SignUpInput struct {
	Email     string          `json:"email" validate:"required,email"`
	Password  string          `json:"password" validate:"required,min=8"`
	Password2 string          `json:"password2" validate:"eqfield=Password"`
	Role      entity.UserRole `json:"role" validate:"required,role"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
