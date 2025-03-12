package dto

type CreateCategoryInput struct {
	Name         string `json:"name" validate:"required,min=5"`
	CoverImg     string `json:"coverImg" validate:"required"`
	Address      string `json:"address" validate:"required"`
	CategoryName string `json:"categoryName" validate:"required"`
}

type EditCategoryInput struct {
	CategoryID   string `json:"categoryId" validate:"required"`
	Name         string `json:"name,omitempty" validate:"min=5"`
	CoverImg     string `json:"coverImg,omitempty"`
	Address      string `json:"address,omitempty"`
	CategoryName string `json:"categoryName,omitempty"`
}

type DeleteCategoryInput struct {
	CategoryID string `json:"categoryId" validate:"required"`
}

type CategorysInput struct {
	Page int `query:"page" validate:"required,min=1"`
}

type CategoryInput struct {
	CategoryID string `json:"categoryId" validate:"required"`
}

type SearchCategoryInput struct {
	Query string `query:"query" validate:"required"`
	Page  int    `query:"page" validate:"required,min=1"`
}
