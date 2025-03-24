package response

import "ubereats/app/core/entity"

type UserResponse struct {
	entity.CoreEntity
	Email string          `json:"email" validate:"email"`
	Role  entity.UserRole `json:"role"`
}

type OwnerResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type GetAllUserOutput struct{}
type CreateAccountOutput struct{}
type LoginOutput struct{}
type MeOutput struct{}
