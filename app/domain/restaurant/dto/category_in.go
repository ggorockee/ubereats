package dto

type CreateCategoryIn struct {
	Name     string `json:"name" validate:"required" mapstructure:"name"`
	CoverImg string `json:"cover_img" validate:"required" mapstructure:"cover_img"`
}
