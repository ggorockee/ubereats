package dto

type CreateCategoryInput struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryInput struct{}
