package response

import "ubereats/app/core/entity"

type UserResponse struct {
	entity.CoreEntity
	Email string          `json:"email" validate:"email"`
	Role  entity.UserRole `json:"role"`
}

func GenUserRes(m *entity.User) UserResponse {
	return UserResponse{
		CoreEntity: entity.CoreEntity{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Email: m.Email,
		Role:  m.Role,
	}
}
