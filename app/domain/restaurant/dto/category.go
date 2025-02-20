package dto

type CategoryInput struct {
	Page int `query:"page" validate:"required,min=1"`
}

type Category struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	CoverImg string `json:"coverImg,omitempty"`
}
